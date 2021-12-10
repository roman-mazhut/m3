// Copyright (c) 2019 Uber Technologies, Inc.
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

package handleroptions

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/m3db/m3/src/query/block"
	"github.com/m3db/m3/src/query/storage"
	"github.com/m3db/m3/src/x/headers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDBResultResponseHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	meta := block.NewResultMetadata()
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 0, len(recorder.Header()))

	recorder = httptest.NewRecorder()
	meta.Exhaustive = false
	ex := headers.LimitHeaderSeriesLimitApplied
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, ex, recorder.Header().Get(headers.LimitHeader))

	recorder = httptest.NewRecorder()
	meta.AddWarning("foo", "bar")
	ex = fmt.Sprintf("%s,%s_%s", headers.LimitHeaderSeriesLimitApplied, "foo", "bar")
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, ex, recorder.Header().Get(headers.LimitHeader))

	recorder = httptest.NewRecorder()
	meta.Exhaustive = true
	ex = "foo_bar"
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, ex, recorder.Header().Get(headers.LimitHeader))

	recorder = httptest.NewRecorder()
	meta = block.NewResultMetadata()
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, &storage.FetchOptions{
		Timeout: 5 * time.Second,
	}))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "5s", recorder.Header().Get(headers.TimeoutHeader))

	recorder = httptest.NewRecorder()
	meta = block.NewResultMetadata()
	meta.WaitedIndex = 3
	meta.WaitedSeriesRead = 42
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "{\"waitedIndex\":3,\"waitedSeriesRead\":42}",
		recorder.Header().Get(headers.WaitedHeader))
}

func TestAddDBResultResponseHeadersFetched(t *testing.T) {
	recorder := httptest.NewRecorder()
	meta := block.NewResultMetadata()
	meta.FetchedSeriesCount = 42
	meta.FetchedMetadataCount = 142
	meta.FetchedResponses = 99
	meta.FetchedBytesEstimate = 1072
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 4, len(recorder.Header()))
	assert.Equal(t, "99", recorder.Header().Get(headers.FetchedResponsesHeader))
	assert.Equal(t, "1072", recorder.Header().Get(headers.FetchedBytesEstimateHeader))
	assert.Equal(t, "42", recorder.Header().Get(headers.FetchedSeriesCount))
	assert.Equal(t, "142", recorder.Header().Get(headers.FetchedMetadataCount))
}

func TestAddDBResultResponseHeadersNamespaces(t *testing.T) {
	recorder := httptest.NewRecorder()
	meta := block.NewResultMetadata()
	meta.Namespaces = []string{}
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 0, len(recorder.Header()))

	recorder = httptest.NewRecorder()
	meta = block.NewResultMetadata()
	meta.Namespaces = []string{"default"}
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "default", recorder.Header().Get(headers.NamespacesHeader))

	recorder = httptest.NewRecorder()
	meta = block.NewResultMetadata()
	meta.Namespaces = []string{"default", "myfavoritens"}
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "default,myfavoritens", recorder.Header().Get(headers.NamespacesHeader))
}

func TestAddDBResultResponseHeadersMetadataByName(t *testing.T) {
	recorder := httptest.NewRecorder()
	meta := block.NewResultMetadata()
	meta.MetadataByName = map[string]*block.ResultMetricMetadata{
		"mymetric": {
			NoSamples:    1,
			WithSamples:  2,
			Aggregated:   3,
			Unaggregated: 4,
		},
	}
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 5, len(recorder.Header()))
	assert.Equal(t, "1", recorder.Header().Get(headers.FetchedSeriesNoSamplesCount))
	assert.Equal(t, "2", recorder.Header().Get(headers.FetchedSeriesWithSamplesCount))
	assert.Equal(t, "3", recorder.Header().Get(headers.FetchedAggregatedSeriesCount))
	assert.Equal(t, "4", recorder.Header().Get(headers.FetchedUnaggregatedSeriesCount))
	assert.Equal(t,
		"{\"mymetric\":{\"NoSamples\":1,\"WithSamples\":2,\"Aggregated\":3,\"Unaggregated\":4}}",
		recorder.Header().Get(headers.MetricStats))

	recorder = httptest.NewRecorder()
	meta = block.NewResultMetadata()
	meta.MetadataByName = map[string]*block.ResultMetricMetadata{
		"metric_a": {
			NoSamples:    1,
			WithSamples:  2,
			Aggregated:   3,
			Unaggregated: 4,
		},
		"metric_b": {
			NoSamples:    10,
			WithSamples:  20,
			Aggregated:   30,
			Unaggregated: 40,
		},
	}
	require.NoError(t, AddDBResultResponseHeaders(recorder, meta, nil))
	assert.Equal(t, 5, len(recorder.Header()))
	assert.Equal(t, "11", recorder.Header().Get(headers.FetchedSeriesNoSamplesCount))
	assert.Equal(t, "22", recorder.Header().Get(headers.FetchedSeriesWithSamplesCount))
	assert.Equal(t, "33", recorder.Header().Get(headers.FetchedAggregatedSeriesCount))
	assert.Equal(t, "44", recorder.Header().Get(headers.FetchedUnaggregatedSeriesCount))
	assert.Equal(t,
		"{\"metric_a\":{\"NoSamples\":1,\"WithSamples\":2,\"Aggregated\":3,\"Unaggregated\":4},"+
			"\"metric_b\":{\"NoSamples\":10,\"WithSamples\":20,\"Aggregated\":30,\"Unaggregated\":40}}",
		recorder.Header().Get(headers.MetricStats))
}

func TestAddReturnedLimitResponseHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	require.NoError(t, AddReturnedLimitResponseHeaders(recorder, &ReturnedDataLimited{
		Series:      3,
		Datapoints:  6,
		TotalSeries: 3,
		Limited:     false,
	}, nil))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "{\"Series\":3,\"Datapoints\":6,\"TotalSeries\":3,\"Limited\":false}",
		recorder.Header().Get(headers.ReturnedDataLimitedHeader))

	recorder = httptest.NewRecorder()
	require.NoError(t, AddReturnedLimitResponseHeaders(recorder, nil, &ReturnedMetadataLimited{
		Results:      3,
		TotalResults: 3,
		Limited:      false,
	}))
	assert.Equal(t, 1, len(recorder.Header()))
	assert.Equal(t, "{\"Results\":3,\"TotalResults\":3,\"Limited\":false}",
		recorder.Header().Get(headers.ReturnedMetadataLimitedHeader))
}
