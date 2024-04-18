// Copyright (c) 2017 Uber Technologies, Inc.
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

package server

import (
	"time"

	"github.com/m3db/m3/src/x/instrument"
	xnet "github.com/m3db/m3/src/x/net"
	"github.com/m3db/m3/src/x/retry"
)

const (
	// By default keepAlives are enabled for TCP connections.
	defaultTCPConnectionKeepAlive = true

	// By default the keep alive period is fairly short for fast
	// breaking of stale connections.
	defaultTCPConnectionKeepAlivePeriod = 10 * time.Second
)

type TLSOptions interface {
	// SetMode sets the tls mode
	SetMode(value TLSMode) TLSOptions
	// Mode returns the tls mode
	Mode() TLSMode

	// SetMutualTLSEnabled sets the mutual tls enabled option
	SetMutualTLSEnabled(value bool) TLSOptions
	// MutualTLSEnabled returns the mutual tls enabled option
	MutualTLSEnabled() bool

	// SetCertFile sets the certificate file path
	SetCertFile(value string) TLSOptions
	// CertFile returns the certificate file path
	CertFile() string

	// SetKeyFile sets the private key file path
	SetKeyFile(value string) TLSOptions
	// KeyFile returns the private key file path
	KeyFile() string

	// SetClientCAFile sets the CA file path
	SetClientCAFile(value string) TLSOptions
	// ClientCAFile returns the CA file path
	ClientCAFile() string
}

// Options provide a set of server options
type Options interface {
	// SetInstrumentOptions sets the instrument options
	SetInstrumentOptions(value instrument.Options) Options

	// InstrumentOptions returns the instrument options
	InstrumentOptions() instrument.Options

	// SetRetryOptions sets the retry options
	SetRetryOptions(value retry.Options) Options

	// RetryOptions returns the retry options
	RetryOptions() retry.Options

	// SetTCPConnectionKeepAlive sets the keep alive state for tcp connections.
	SetTCPConnectionKeepAlive(value bool) Options

	// TCPConnectionKeepAlive returns the keep alive state for tcp connections.
	TCPConnectionKeepAlive() bool

	// SetTCPConnectionKeepAlivePeriod sets the keep alive period for tcp connections.
	// NB(xichen): on Linux this modifies both the idle time (i.e,. the time when the
	// last packet is sent from the client and when the first keepAlive probe is sent)
	// as well as the interval between keepAlive probes.
	SetTCPConnectionKeepAlivePeriod(value time.Duration) Options

	// TCPConnectionKeepAlivePeriod returns the keep alive period for tcp connections.
	TCPConnectionKeepAlivePeriod() time.Duration

	// SetListenerOptions sets the listener options for the server.
	SetListenerOptions(value xnet.ListenerOptions) Options

	// ListenerOptions sets the listener options for the server.
	ListenerOptions() xnet.ListenerOptions

	// SetTLSOptions sets the tls options for the server
	SetTLSOptions(value TLSOptions) Options

	// TLSOptions returns the tls options for the server
	TLSOptions() TLSOptions
}

type options struct {
	instrumentOpts               instrument.Options
	retryOpts                    retry.Options
	tcpConnectionKeepAlive       bool
	tcpConnectionKeepAlivePeriod time.Duration
	listenerOpts                 xnet.ListenerOptions
	tlsOptions                   TLSOptions
}

// NewOptions creates a new set of server options
func NewOptions() Options {
	return &options{
		instrumentOpts:               instrument.NewOptions(),
		retryOpts:                    retry.NewOptions(),
		tcpConnectionKeepAlive:       defaultTCPConnectionKeepAlive,
		tcpConnectionKeepAlivePeriod: defaultTCPConnectionKeepAlivePeriod,
		listenerOpts:                 xnet.NewListenerOptions(),
		tlsOptions:                   NewTLSOptions(),
	}
}

func (o *options) SetInstrumentOptions(value instrument.Options) Options {
	opts := *o
	opts.instrumentOpts = value
	return &opts
}

func (o *options) InstrumentOptions() instrument.Options {
	return o.instrumentOpts
}

func (o *options) SetRetryOptions(value retry.Options) Options {
	opts := *o
	opts.retryOpts = value
	return &opts
}

func (o *options) RetryOptions() retry.Options {
	return o.retryOpts
}

func (o *options) SetTCPConnectionKeepAlive(value bool) Options {
	opts := *o
	opts.tcpConnectionKeepAlive = value
	return &opts
}

func (o *options) TCPConnectionKeepAlive() bool {
	return o.tcpConnectionKeepAlive
}

func (o *options) SetTCPConnectionKeepAlivePeriod(value time.Duration) Options {
	opts := *o
	opts.tcpConnectionKeepAlivePeriod = value
	return &opts
}

func (o *options) TCPConnectionKeepAlivePeriod() time.Duration {
	return o.tcpConnectionKeepAlivePeriod
}

func (o *options) SetListenerOptions(value xnet.ListenerOptions) Options {
	opts := *o
	opts.listenerOpts = value
	return &opts
}

func (o *options) ListenerOptions() xnet.ListenerOptions {
	return o.listenerOpts
}

func (o *options) SetTLSOptions(value TLSOptions) Options {
	opts := *o
	opts.tlsOptions = value
	return &opts
}

func (o *options) TLSOptions() TLSOptions {
	return o.tlsOptions
}

type tlsOptions struct {
	mode         TLSMode
	mTLSEnabled  bool
	certFile     string
	keyFile      string
	clientCAFile string
}

// NewTLSOptions creates a new set of tls options
func NewTLSOptions() TLSOptions {
	return &tlsOptions{
		mode:        TLSDisabled,
		mTLSEnabled: false,
	}
}

func (o *tlsOptions) SetMode(value TLSMode) TLSOptions {
	opts := *o
	opts.mode = value
	return &opts
}

func (o *tlsOptions) Mode() TLSMode {
	return o.mode
}

func (o *tlsOptions) SetMutualTLSEnabled(value bool) TLSOptions {
	opts := *o
	opts.mTLSEnabled = value
	return &opts
}

func (o *tlsOptions) MutualTLSEnabled() bool {
	return o.mTLSEnabled
}

func (o *tlsOptions) SetCertFile(value string) TLSOptions {
	opts := *o
	opts.certFile = value
	return &opts
}

func (o *tlsOptions) CertFile() string {
	return o.certFile
}

func (o *tlsOptions) SetKeyFile(value string) TLSOptions {
	opts := *o
	opts.keyFile = value
	return &opts
}

func (o *tlsOptions) KeyFile() string {
	return o.keyFile
}

func (o *tlsOptions) SetClientCAFile(value string) TLSOptions {
	opts := *o
	opts.clientCAFile = value
	return &opts
}

func (o *tlsOptions) ClientCAFile() string {
	return o.clientCAFile
}
