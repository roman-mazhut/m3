// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package aggregator

import (
	"fmt"
	"math"
	"sync"
	"time"

	raggregation "github.com/m3db/m3/src/aggregator/aggregation"
	"github.com/m3db/m3/src/metrics/metadata"
	"github.com/m3db/m3/src/metrics/metric/aggregated"
	"github.com/m3db/m3/src/metrics/metric/unaggregated"
	"github.com/m3db/m3/src/metrics/transformation"
	"github.com/m3db/m3/src/x/instrument"
	xtime "github.com/m3db/m3/src/x/time"

	"github.com/willf/bitset"
	"go.uber.org/zap"
)

type lockedCounterAggregation struct {
	aggregation counterAggregation
	sourcesSeen map[uint32]*bitset.BitSet
	mtx         sync.Mutex
	dirty       bool
	// resendEnabled is allowed to change while an aggregation is open, so it must be behind the lock.
	resendEnabled bool
	closed        bool
	lastUpdatedAt xtime.UnixNano
}

type timedCounter struct {
	lockedAgg  *lockedCounterAggregation
	startAt    xtime.UnixNano // start time of an aggregation window
	prevStart  xtime.UnixNano
	nextStart  xtime.UnixNano
	inDirtySet bool
}

// close is called when the aggregation has been expired or the element is being closed.
func (ta *timedCounter) close() {
	ta.lockedAgg.aggregation.Close()
	ta.lockedAgg = nil
}

// CounterElem is an element storing time-bucketed aggregations.
type CounterElem struct {
	counterElemBase
	elemBase
	// startTime -> agg (new one per every resolution)
	values map[xtime.UnixNano]timedCounter
	// startTime -> state. this is local state to the flusher and does not need to guarded with a lock.
	// values and flushState should always have the exact same key set.
	flushState map[xtime.UnixNano]flushState
	// sorted start aligned times that have been written to since the last flush
	dirty []xtime.UnixNano

	// internal/no need for synchronization: small buffers to avoid memory allocations during consumption
	toConsume            []consumeState
	flushStateToExpire   []xtime.UnixNano
	forwardTimesToExpire []xtime.UnixNano
	// end internal state

	// min time in the values map. allows for iterating through map.
	minStartTime xtime.UnixNano
	// max time in the values map. allows for iterating through map.
	maxStartTime xtime.UnixNano
}

// NewCounterElem returns a new CounterElem.
func NewCounterElem(data ElemData, opts ElemOptions) (*CounterElem, error) {
	e := &CounterElem{
		elemBase:   newElemBase(opts),
		dirty:      make([]xtime.UnixNano, 0, defaultNumAggregations), // in most cases values will have two entries
		values:     make(map[xtime.UnixNano]timedCounter),
		flushState: make(map[xtime.UnixNano]flushState),
	}
	if err := e.ResetSetData(data); err != nil {
		return nil, err
	}
	return e, nil
}

// MustNewCounterElem returns a new CounterElem and panics if an error occurs.
func MustNewCounterElem(data ElemData, opts ElemOptions) *CounterElem {
	elem, err := NewCounterElem(data, opts)
	if err != nil {
		panic(fmt.Errorf("unable to create element: %v", err))
	}
	return elem
}

// ResetSetData resets the element and sets data.
func (e *CounterElem) ResetSetData(data ElemData) error {
	useDefaultAggregation := data.AggTypes.IsDefault()
	if useDefaultAggregation {
		data.AggTypes = e.DefaultAggregationTypes(e.aggTypesOpts)
	}
	if err := e.elemBase.resetSetData(data, useDefaultAggregation); err != nil {
		return err
	}
	return e.counterElemBase.ResetSetData(e.aggTypesOpts, data.AggTypes, useDefaultAggregation)
}

// AddUnion adds a metric value union at a given timestamp.
func (e *CounterElem) AddUnion(timestamp time.Time, mu unaggregated.MetricUnion, resendEnabled bool) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window)
	lockedAgg, err := e.findOrCreate(alignedStart.UnixNano(), createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.mtx.Lock()
	if lockedAgg.closed {
		// Note: this might have created an entry in the dirty set for lockedAgg when calling findOrCreate, even though
		// it's already closed. The Consume loop will detect this and clean it up.
		aggResendEnabled := lockedAgg.resendEnabled
		lockedAgg.mtx.Unlock()
		if !aggResendEnabled && resendEnabled {
			return errClosedBeforeResendEnabledMigration
		}
		return errAggregationClosed
	}
	lockedAgg.aggregation.AddUnion(timestamp, mu)
	lockedAgg.dirty = true
	lockedAgg.lastUpdatedAt = xtime.Now()
	lockedAgg.resendEnabled = resendEnabled
	lockedAgg.mtx.Unlock()
	return nil
}

// AddValue adds a metric value at a given timestamp.
func (e *CounterElem) AddValue(timestamp time.Time, value float64, annotation []byte) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.mtx.Lock()
	if lockedAgg.closed {
		lockedAgg.mtx.Unlock()
		return errAggregationClosed
	}
	lockedAgg.aggregation.Add(timestamp, value, annotation)
	lockedAgg.dirty = true
	lockedAgg.lastUpdatedAt = xtime.Now()
	lockedAgg.mtx.Unlock()
	return nil
}

// AddUnique adds a metric value from a given source at a given timestamp.
// If previous values from the same source have already been added to the
// same aggregation, the incoming value is discarded.
//nolint: dupl
func (e *CounterElem) AddUnique(
	timestamp time.Time,
	metric aggregated.ForwardedMetric,
	metadata metadata.ForwardMetadata,
) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{
		initSourceSet: true,
	})
	if err != nil {
		return err
	}
	lockedAgg.mtx.Lock()
	if lockedAgg.closed {
		lockedAgg.mtx.Unlock()
		return errAggregationClosed
	}
	versionsSeen := lockedAgg.sourcesSeen[metadata.SourceID]
	if versionsSeen == nil {
		// N.B - these bitsets will be transitively cached through the cached sources seen.
		versionsSeen = bitset.New(defaultNumVersions)
		lockedAgg.sourcesSeen[metadata.SourceID] = versionsSeen
	}
	version := uint(metric.Version)
	if versionsSeen.Test(version) {
		lockedAgg.mtx.Unlock()
		return errDuplicateForwardingSource
	}
	versionsSeen.Set(version)

	if metric.Version > 0 {
		e.writeMetrics.updatedValues.Inc(1)
		for i := range metric.Values {
			if err := lockedAgg.aggregation.UpdateVal(timestamp, metric.Values[i], metric.PrevValues[i]); err != nil {
				return err
			}
		}
	} else {
		for _, v := range metric.Values {
			lockedAgg.aggregation.Add(timestamp, v, metric.Annotation)
		}
	}
	lockedAgg.dirty = true
	lockedAgg.lastUpdatedAt = xtime.Now()
	lockedAgg.resendEnabled = metadata.ResendEnabled
	lockedAgg.mtx.Unlock()
	return nil
}

// remove expired aggregations from the values map.
func (e *CounterElem) expireValuesWithLock(
	targetNanos int64,
	isEarlierThanFn isEarlierThanFn,
	flushMetrics flushMetrics) {
	e.flushStateToExpire = e.flushStateToExpire[:0]
	if len(e.values) == 0 {
		return
	}
	resolution := e.sp.Resolution().Window

	currAgg := e.values[e.minStartTime]
	resendExpire := targetNanos - int64(e.bufferForPastTimedMetricFn(resolution))
	for isEarlierThanFn(int64(currAgg.startAt), resolution, targetNanos) {
		if e.flushState[currAgg.startAt].latestResendEnabled {
			// if resend enabled we want to keep this value until it is outside the buffer past period.
			if !isEarlierThanFn(int64(currAgg.startAt), resolution, resendExpire) {
				break
			}
		}

		// close the agg to prevent any more writes.
		dirty := false
		currAgg.lockedAgg.mtx.Lock()
		if currAgg.lockedAgg.resendEnabled != e.flushState[currAgg.startAt].latestResendEnabled {
			// the aggregation migrated to resendEnabled after the flusher read the resendEnabled state.
			// keep the aggregation for now and try to expire on the next flush.
			currAgg.lockedAgg.mtx.Unlock()
			break
		}
		currAgg.lockedAgg.closed = true
		dirty = currAgg.lockedAgg.dirty
		currAgg.lockedAgg.mtx.Unlock()
		if dirty {
			// a race occurred and a write happened before we could close the aggregation. will expire next time.
			break
		}

		// if this current value is closed and clean it will no longer be flushed. this means it's safe
		// to remove the previous value since it will no longer be needed for binary transformations. when the
		// next value is eligible to be expired, this current value will actually be removed.
		// if we're currently pointing at the start skip this because there is no previous for the start. this
		// ensures we always keep at least one value in the map for binary transformations.
		if prevAgg, ok := e.prevAggWithLock(currAgg); ok && currAgg.startAt != e.minStartTime {
			// can't expire flush state until after the flushing, so we save the time to expire later.
			e.flushStateToExpire = append(e.flushStateToExpire, e.minStartTime)
			delete(e.values, e.minStartTime)
			e.minStartTime = currAgg.startAt
			flushMetrics.valuesExpired.Inc(1)

			// it's safe to access this outside the agg lock since it was closed in a previous iteration.
			// This is to make sure there aren't too many cached source sets taking up
			// too much space.
			if prevAgg.lockedAgg.sourcesSeen != nil && len(e.cachedSourceSets) < e.opts.MaxNumCachedSourceSets() {
				e.cachedSourceSets = append(e.cachedSourceSets, prevAgg.lockedAgg.sourcesSeen)
			}
			prevAgg.close()
		}
		var ok bool
		currAgg, ok = e.nextAggWithLock(currAgg)
		if !ok {
			break
		}
	}
}

func (e *CounterElem) expireFlushState() {
	for _, t := range e.flushStateToExpire {
		fState, ok := e.flushState[t]
		if !ok {
			ts := t.ToTime()
			instrument.EmitAndLogInvariantViolation(e.opts.InstrumentOptions(), func(l *zap.Logger) {
				l.Error("expire time not in state map", zap.Time("ts", ts))
			})
			continue
		}
		fState.close()
		delete(e.flushState, t)
	}
}

// return the previous aggregation before the provided time. returns false if the provided time is the
// earliest time or the map is empty.
func (e *CounterElem) prevAggWithLock(agg timedCounter) (timedCounter, bool) {
	if len(e.values) == 0 {
		return timedCounter{}, false
	}
	if agg.prevStart != 0 {
		prevAgg, ok := e.values[agg.prevStart]
		return prevAgg, ok
	}

	resolution := e.sp.Resolution().Window
	startTime := agg.startAt.Add(-resolution)
	for !startTime.Before(e.minStartTime) {
		agg, ok := e.values[startTime]
		if ok {
			return agg, true
		}
		startTime = startTime.Add(-resolution)
	}
	return timedCounter{}, false
}

// return the next aggregation after the provided time. returns false if the provided time is the
// largest time or the map is empty.
func (e *CounterElem) nextAggWithLock(agg timedCounter) (timedCounter, bool) {
	if len(e.values) == 0 {
		return timedCounter{}, false
	}
	if agg.nextStart != 0 {
		nextAgg, ok := e.values[agg.nextStart]
		return nextAgg, ok
	}
	resolution := e.sp.Resolution().Window
	start := agg.startAt.Add(resolution)
	for !start.After(e.maxStartTime) {
		agg, ok := e.values[start]
		if ok {
			return agg, true
		}
		start = start.Add(resolution)
	}
	return timedCounter{}, false
}

// Consume consumes values before a given time and removes them from the element
// after they are consumed, returning whether the element can be collected after
// the consumption is completed.
// NB: Consume is not thread-safe and must be called within a single goroutine
// to avoid race conditions.
func (e *CounterElem) Consume(
	targetNanos int64,
	isEarlierThanFn isEarlierThanFn,
	timestampNanosFn timestampNanosFn,
	targetNanosFn targetNanosFn,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
	onForwardedFlushedFn onForwardingElemFlushedFn,
	jitter time.Duration,
	flushType flushType,
) bool {
	resolution := e.sp.Resolution().Window
	fMetrics := e.flushMetrics(resolution, flushType)
	fMetrics.valuesProcessed.Inc(1)

	// reverse engineer the allowed lateness.
	latenessAllowed := time.Duration(targetNanos - targetNanosFn(targetNanos))
	e.Lock()
	if e.closed {
		e.Unlock()
		return false
	}

	// move currently dirty aggs to toConsume to process next.
	e.dirtyToConsumeWithLock(targetNanos, resolution, isEarlierThanFn)

	// expire the values and aggregations while we still hold the lock.
	e.expireValuesWithLock(targetNanos, isEarlierThanFn, fMetrics)
	canCollect := len(e.dirty) == 0 && e.tombstoned
	e.Unlock()

	// Process the aggregations that are ready for consumption.
	for _, cState := range e.toConsume {
		e.processValue(cState,
			timestampNanosFn,
			flushLocalFn,
			flushForwardedFn,
			resolution,
			latenessAllowed,
			jitter,
			fMetrics,
		)
	}

	// expire the flush state after processing since it's needed in the processing.
	e.expireFlushState()

	if e.parsedPipeline.HasRollup {
		forwardedAggregationKey, _ := e.ForwardedAggregationKey()
		e.forwardTimesToExpire = e.forwardTimesToExpire[:0]
		for _, startTime := range e.flushStateToExpire {
			// the forward writer uses the timestamp of the aggregation, so need to convert the start aligned time
			// to a timestamp.
			e.forwardTimesToExpire = append(e.forwardTimesToExpire,
				xtime.UnixNano(timestampNanosFn(int64(startTime), resolution)))
		}
		onForwardedFlushedFn(e.onForwardedAggregationWrittenFn, forwardedAggregationKey, e.forwardTimesToExpire)
	}

	return canCollect
}

func (e *CounterElem) dirtyToConsumeWithLock(targetNanos int64,
	resolution time.Duration,
	isEarlierThanFn isEarlierThanFn) {
	e.toConsume = e.toConsume[:0]
	// Evaluate and GC expired items.
	dirtyTimes := e.dirty
	e.dirty = e.dirty[:0]
	for i, dirtyTime := range dirtyTimes {
		if !isEarlierThanFn(int64(dirtyTime), resolution, targetNanos) {
			// not ready yet
			e.dirty = append(e.dirty, dirtyTime)
			continue
		}
		agg, ok := e.values[dirtyTime]
		if !ok {
			// there is a race where a writer adds a closed aggregation to the dirty set. eventually the closed
			// aggregation is expired and removed from the values map. ok to skip.
			continue
		}

		var dirty bool
		e.toConsume, dirty = e.appendConsumeStateWithLock(agg, e.toConsume, isDirty)
		if !dirty {
			// there is a race where the value was added to the dirty set, but the writer didn't actually update the
			// value yet (by marking dirty). add back to the dirty set so it can be processed in the next round once
			// the value has been updated.
			e.dirty = append(e.dirty, dirtyTime)
			continue
		}
		val := e.values[dirtyTime]
		val.inDirtySet = false
		e.values[dirtyTime] = val
		cState := e.toConsume[len(e.toConsume)-1]

		// potentially consume the nextAgg as well in case we need to cascade an update to the nextAgg.
		// this is necessary for binary transformations that rely on the previous aggregation value for calculating the
		// current aggregation value. if the nextAgg was already flushed, it used an outdated value for the previous
		// value (this agg). this can only happen when we allow updating previously flushed data (i.e resendEnabled).
		if cState.resendEnabled {
			nextAgg, ok := e.nextAggWithLock(agg)
			// only need to add if not already in the dirty set (since it will be added in a subsequent iteration).
			if ok &&
				// at the end of the dirty times OR the next dirty time does not match.
				(i == len(dirtyTimes)-1 || dirtyTimes[i+1] != nextAgg.startAt) {
				// only need to add if it was previously flushed.
				e.toConsume, _ = e.appendConsumeStateWithLock(nextAgg, e.toConsume, e.isFlushed)
			}
		}
	}
}

func (e *CounterElem) isFlushed(c consumeState) bool {
	return e.flushState[c.startAt].flushed
}

// append the consumeState for the timedCounter to the provided slice if it matches the provided filter.
// returns the updated slice and true if added.
func (e *CounterElem) appendConsumeStateWithLock(
	agg timedCounter,
	toConsume []consumeState,
	includeFilter func(consumeState) bool) ([]consumeState, bool) {
	// eagerly append a new element so we can try reusing memory already allocated in the slice.
	toConsume = append(toConsume, consumeState{})
	cState := toConsume[len(toConsume)-1]
	if cState.values == nil {
		cState.values = make([]float64, len(e.aggTypes))
	}
	cState.values = cState.values[:0]
	// copy the lockedAgg data while holding the lock.
	agg.lockedAgg.mtx.Lock()
	cState.dirty = agg.lockedAgg.dirty
	cState.lastUpdatedAt = agg.lockedAgg.lastUpdatedAt
	cState.resendEnabled = agg.lockedAgg.resendEnabled
	for _, aggType := range e.aggTypes {
		cState.values = append(cState.values, agg.lockedAgg.aggregation.ValueOf(aggType))
	}
	cState.annotation = raggregation.MaybeReplaceAnnotation(
		cState.annotation, agg.lockedAgg.aggregation.Annotation())
	agg.lockedAgg.dirty = false
	agg.lockedAgg.mtx.Unlock()

	// update with everything else.
	prevAgg, ok := e.prevAggWithLock(agg)
	if ok {
		cState.prevStartTime = prevAgg.startAt
	} else {
		cState.prevStartTime = 0
	}
	cState.startAt = agg.startAt
	toConsume[len(toConsume)-1] = cState
	// update the flush state with the latestResendEnabled since expireValuesWithLock needs it before actual processing.
	fState := e.flushState[cState.startAt]
	fState.latestResendEnabled = cState.resendEnabled
	e.flushState[cState.startAt] = fState

	if includeFilter != nil && !includeFilter(cState) {
		// since we eagerly appended, we need to remove if it should not be included.
		toConsume = toConsume[0 : len(toConsume)-1]
		return toConsume, false
	}
	return toConsume, true
}

// Close closes the element.
func (e *CounterElem) Close() {
	e.Lock()
	if e.closed {
		e.Unlock()
		return
	}
	e.closed = true
	e.id = nil
	e.parsedPipeline = parsedPipeline{}
	e.writeForwardedMetricFn = nil
	e.onForwardedAggregationWrittenFn = nil
	for idx := range e.cachedSourceSets {
		e.cachedSourceSets[idx] = nil
	}
	e.cachedSourceSets = nil

	// note: this is not in the hot path so it's ok to iterate over the map.
	// this allows to catch any bugs with unexpected entries still in the map.
	minStartTime := e.minStartTime
	for k, v := range e.values {
		if k < minStartTime {
			k := k
			ts := e.minStartTime.ToTime()
			instrument.EmitAndLogInvariantViolation(e.opts.InstrumentOptions(), func(l *zap.Logger) {
				l.Error("value timestamp is less than min start time",
					zap.Time("ts", k.ToTime()),
					zap.Time("min", ts))
			})
		}
		v.close()
		delete(e.values, k)
		fState, ok := e.flushState[k]
		if ok {
			fState.close()
		}
		delete(e.flushState, k)
	}
	// clean up any dangling flush state that should never exist.
	for k, v := range e.flushState {
		ts := k.ToTime()
		instrument.EmitAndLogInvariantViolation(e.opts.InstrumentOptions(), func(l *zap.Logger) {
			l.Error("dangling state timestamp", zap.Time("ts", ts))
		})
		v.close()
		delete(e.flushState, k)
	}
	e.counterElemBase.Close()
	aggTypesPool := e.aggTypesOpts.TypesPool()
	pool := e.ElemPool(e.opts)
	e.dirty = e.dirty[:0]
	e.toConsume = e.toConsume[:0]
	e.flushStateToExpire = e.flushStateToExpire[:0]
	e.minStartTime = 0
	e.Unlock()

	if !e.useDefaultAggregation {
		aggTypesPool.Put(e.aggTypes)
	}
	pool.Put(e)
}

func (e *CounterElem) insertDirty(alignedStart xtime.UnixNano) {
	numValues := len(e.dirty)

	// Optimize for the common case.
	if numValues > 0 && e.dirty[numValues-1] == alignedStart {
		return
	}
	// Binary search for the unusual case. We intentionally do not
	// use the sort.Search() function because it requires passing
	// in a closure.
	left, right := 0, numValues
	for left < right {
		mid := left + (right-left)/2 // avoid overflow
		if e.dirty[mid] < alignedStart {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// If the current timestamp is equal to or larger than the target time,
	// return the index as is.
	if left < numValues && e.dirty[left] == alignedStart {
		return
	}

	e.dirty = append(e.dirty, 0)
	copy(e.dirty[left+1:numValues+1], e.dirty[left:numValues])
	e.dirty[left] = alignedStart
}

// find finds the aggregation for a given time, or returns nil.
//nolint: dupl
func (e *CounterElem) find(alignedStartNanos xtime.UnixNano) (timedCounter, error) {
	e.RLock()
	if e.closed {
		e.RUnlock()
		return timedCounter{}, errElemClosed
	}
	timedAgg, ok := e.values[alignedStartNanos]
	if ok {
		e.RUnlock()
		return timedAgg, nil
	}
	e.RUnlock()
	return timedCounter{}, nil
}

// findOrCreate finds the aggregation for a given time, or creates one
// if it doesn't exist.
//nolint: dupl
func (e *CounterElem) findOrCreate(
	alignedStartNanos int64,
	createOpts createAggregationOptions,
) (*lockedCounterAggregation, error) {
	e.writeMetrics.writes.Inc(1)
	alignedStart := xtime.UnixNano(alignedStartNanos)
	found, err := e.find(alignedStart)
	if err != nil {
		return nil, err
	}
	// if the aggregation is found and does not need to be updated, return as is.
	if found.lockedAgg != nil && found.inDirtySet {
		return found.lockedAgg, err
	}

	e.Lock()
	if e.closed {
		e.Unlock()
		return nil, errElemClosed
	}

	timedAgg, ok := e.values[alignedStart]
	if ok {
		// add to dirty set so it will be flushed.
		if !timedAgg.inDirtySet {
			timedAgg.inDirtySet = true
			e.insertDirty(alignedStart)
			e.values[alignedStart] = timedAgg
		}
		e.Unlock()
		return timedAgg.lockedAgg, nil
	}

	var sourcesSeen map[uint32]*bitset.BitSet
	if createOpts.initSourceSet {
		if numCachedSourceSets := len(e.cachedSourceSets); numCachedSourceSets > 0 {
			sourcesSeen = e.cachedSourceSets[numCachedSourceSets-1]
			e.cachedSourceSets[numCachedSourceSets-1] = nil
			e.cachedSourceSets = e.cachedSourceSets[:numCachedSourceSets-1]
			for _, bs := range sourcesSeen {
				bs.ClearAll()
			}
		} else {
			sourcesSeen = make(map[uint32]*bitset.BitSet)
		}
	}
	timedAgg = timedCounter{
		startAt: alignedStart,
		lockedAgg: &lockedCounterAggregation{
			sourcesSeen: sourcesSeen,
			aggregation: e.NewAggregation(e.opts, e.aggOpts),
		},
		inDirtySet: true,
	}

	if len(e.values) == 0 || e.minStartTime > alignedStart {
		e.minStartTime = alignedStart
	}
	prevMaxStart := e.maxStartTime
	if len(e.values) == 0 || alignedStart > e.maxStartTime {
		e.maxStartTime = alignedStart
	}

	if len(e.values) > 0 {
		if e.maxStartTime == alignedStart {
			// common case we are adding the latest start time.
			timedAgg.prevStart = prevMaxStart
			prevAgg := e.values[prevMaxStart]
			prevAgg.nextStart = alignedStart
			e.values[prevMaxStart] = prevAgg
		} else {
			// look up
			prevAgg, ok := e.prevAggWithLock(timedAgg)
			if ok {
				timedAgg.prevStart = prevAgg.startAt
				prevAgg.nextStart = alignedStart
				e.values[prevAgg.startAt] = prevAgg
			}
			nextAgg, ok := e.nextAggWithLock(timedAgg)
			if ok {
				timedAgg.nextStart = nextAgg.startAt
				nextAgg.prevStart = alignedStart
				e.values[nextAgg.startAt] = nextAgg
			}
		}
	}

	e.values[alignedStart] = timedAgg
	e.insertDirty(alignedStart)
	e.Unlock()
	return timedAgg.lockedAgg, nil
}

// returns true if a datapoint is emitted.
func (e *CounterElem) processValue(
	cState consumeState,
	timestampNanosFn timestampNanosFn,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
	resolution time.Duration,
	latenessAllowed time.Duration,
	jitter time.Duration,
	flushMetrics flushMetrics) {
	var (
		transformations  = e.parsedPipeline.Transformations
		discardNaNValues = e.opts.DiscardNaNAggregatedValues()
		timestamp        = xtime.UnixNano(timestampNanosFn(int64(cState.startAt), resolution))
		prevTimestamp    = xtime.UnixNano(timestampNanosFn(int64(cState.prevStartTime), resolution))
		// expectedProcessingTime should be the next resolution window after the aggregation was updated.
		expectedProcessingTime = cState.lastUpdatedAt.Truncate(resolution).Add(resolution)
	)
	fState := e.flushState[cState.startAt]
	if cState.dirty && fState.flushed && !cState.resendEnabled {
		cState := cState
		instrument.EmitAndLogInvariantViolation(e.opts.InstrumentOptions(), func(l *zap.Logger) {
			l.Error("reflushing aggregation without resendEnabled", zap.Any("consumeState", cState))
		})
	}
	flushMetrics.valuesProcessed.Inc(1)
	for aggTypeIdx, aggType := range e.aggTypes {
		var extraDp transformation.Datapoint
		value := cState.values[aggTypeIdx]
		for _, transformOp := range transformations {
			unaryOp, isUnaryOp := transformOp.UnaryTransform()
			binaryOp, isBinaryOp := transformOp.BinaryTransform()
			unaryMultiOp, isUnaryMultiOp := transformOp.UnaryMultiOutputTransform()
			switch {
			case isUnaryOp:
				curr := transformation.Datapoint{
					TimeNanos: int64(timestamp),
					Value:     value,
				}

				res := unaryOp.Evaluate(curr)

				value = res.Value

			case isBinaryOp:
				prev := transformation.Datapoint{
					Value: nan,
				}
				if cState.prevStartTime > 0 {
					prevFlushState, ok := e.flushState[cState.prevStartTime]
					if !ok {
						ts := cState.prevStartTime.ToTime()
						instrument.EmitAndLogInvariantViolation(e.opts.InstrumentOptions(), func(l *zap.Logger) {
							l.Error("previous start time not in state map",
								zap.Time("ts", ts))
						})
					} else {
						prev.Value = prevFlushState.consumedValues[aggTypeIdx]
						prev.TimeNanos = int64(prevTimestamp)
					}
				}
				curr := transformation.Datapoint{
					TimeNanos: int64(timestamp),
					Value:     value,
				}
				res := binaryOp.Evaluate(prev, curr, transformation.FeatureFlags{})

				// NB: we only need to record the value needed for derivative transformations.
				// We currently only support first-order derivative transformations so we only
				// need to keep one value. In the future if we need to support higher-order
				// derivative transformations, we need to store an array of values here.
				if fState.consumedValues == nil {
					fState.consumedValues = make([]float64, len(e.aggTypes))
				}
				fState.consumedValues[aggTypeIdx] = curr.Value
				value = res.Value
			case isUnaryMultiOp:
				curr := transformation.Datapoint{
					TimeNanos: int64(timestamp),
					Value:     value,
				}

				var res transformation.Datapoint
				res, extraDp = unaryMultiOp.Evaluate(curr, resolution)
				value = res.Value
			}
		}

		if discardNaNValues && math.IsNaN(value) {
			continue
		}

		// It's ok to send a 0 prevValue on the first forward because it's not used in AddUnique unless it's a
		// resend (version > 0)
		var prevValue float64
		if fState.emittedValues == nil {
			fState.emittedValues = make([]float64, len(e.aggTypes))
		} else {
			prevValue = fState.emittedValues[aggTypeIdx]
		}
		fState.emittedValues[aggTypeIdx] = value
		if fState.flushed {
			// no need to resend a value that hasn't changed.
			if (math.IsNaN(prevValue) && math.IsNaN(value)) || (prevValue == value) {
				continue
			}
		}

		fwdType := forwardTypeRemote
		if !e.parsedPipeline.HasRollup {
			fwdType = forwardTypeLocal
			toFlush := make([]transformation.Datapoint, 0, 2)
			toFlush = append(toFlush, transformation.Datapoint{
				TimeNanos: int64(timestamp),
				Value:     value,
			})
			if extraDp.TimeNanos != 0 {
				toFlush = append(toFlush, extraDp)
			}
			for _, point := range toFlush {
				switch e.idPrefixSuffixType {
				case NoPrefixNoSuffix:
					flushLocalFn(nil, e.id, nil, point.TimeNanos, point.Value, cState.annotation,
						e.sp)
				case WithPrefixWithSuffix:
					flushLocalFn(e.FullPrefix(e.opts), e.id, e.TypeStringFor(e.aggTypesOpts, aggType),
						point.TimeNanos, point.Value, cState.annotation, e.sp)
				}
			}
		} else {
			forwardedAggregationKey, _ := e.ForwardedAggregationKey()
			flushForwardedFn(e.writeForwardedMetricFn, forwardedAggregationKey,
				int64(timestamp), value, prevValue, cState.annotation, cState.resendEnabled)
		}
		// add latenessAllowed and jitter to the timestamp of the aggregation, since those should not be
		// counted towards the processing lag.
		// forward lag = current time - (agg timestamp + lateness allowed + jitter)
		// use expectedProcessingTime instead of the aggregation timestamp since the aggregation timestamp could be
		// in the past for updated aggregations (resendEnabled).
		flushMetrics.forwardLag(forwardKey{fwdType: fwdType, jitter: false}).
			RecordDuration(xtime.Since(expectedProcessingTime.Add(latenessAllowed + jitter)))
		flushMetrics.forwardLag(forwardKey{fwdType: fwdType, jitter: true}).
			RecordDuration(xtime.Since(expectedProcessingTime.Add(latenessAllowed)))
	}
	fState.flushed = true
	e.flushState[cState.startAt] = fState
}
