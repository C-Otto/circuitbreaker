package main

import (
	"context"
	"testing"
	"time"

	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	defer Timeout()()

	t.Run("settle", func(t *testing.T) {
		testProcess(t, resolveEventSettle)
	})
	t.Run("forward fail", func(t *testing.T) {
		testProcess(t, resolveEventForwardFail)
	})
	t.Run("link fail", func(t *testing.T) {
		testProcess(t, resolveEventLinkFail)
	})
}

type resolveEvent int

const (
	resolveEventSettle resolveEvent = iota
	resolveEventForwardFail
	resolveEventLinkFail
)

func testProcess(t *testing.T, event resolveEvent) {
	cfg := &config{
		groupConfig: groupConfig{
			MaxPendingHtlcs: 2,
		},
	}

	client := newLndclientMock()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newProcess(client, cfg)

	resolved := make(chan struct{})
	p.resolvedCallback = func() {
		close(resolved)
	}

	exit := make(chan error)
	go func() {
		exit <- p.run(ctx)
	}()

	key := &routerrpc.CircuitKey{
		ChanId: 2,
		HtlcId: 5,
	}
	client.htlcInterceptorRequests <- &routerrpc.ForwardHtlcInterceptRequest{
		IncomingCircuitKey: key,
	}

	resp := <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_RESUME, resp.Action)

	htlcEvent := &routerrpc.HtlcEvent{
		EventType:         routerrpc.HtlcEvent_FORWARD,
		IncomingChannelId: key.ChanId,
		IncomingHtlcId:    key.HtlcId,
	}

	switch event {
	case resolveEventForwardFail:
		htlcEvent.Event = &routerrpc.HtlcEvent_ForwardFailEvent{}

	case resolveEventLinkFail:
		htlcEvent.Event = &routerrpc.HtlcEvent_LinkFailEvent{}

	case resolveEventSettle:
		htlcEvent.Event = &routerrpc.HtlcEvent_SettleEvent{}
	}

	client.htlcEvents <- htlcEvent

	<-resolved

	cancel()
	require.ErrorIs(t, <-exit, context.Canceled)
}

func TestRateLimit(t *testing.T) {
	defer Timeout()()

	cfg := &config{
		groupConfig: groupConfig{
			HtlcMinInterval: time.Minute,
			HtlcBurstSize:   2,
		},
	}

	client := newLndclientMock()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newProcess(client, cfg)

	exit := make(chan error)
	go func() {
		exit <- p.run(ctx)
	}()

	key := &routerrpc.CircuitKey{
		ChanId: 2,
		HtlcId: 5,
	}
	interceptReq := &routerrpc.ForwardHtlcInterceptRequest{
		IncomingCircuitKey: key,
	}

	// First htlc accepted.
	client.htlcInterceptorRequests <- interceptReq
	resp := <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_RESUME, resp.Action)

	// Second htlc right after is also accepted because of burst size 2.
	interceptReq.IncomingCircuitKey.HtlcId++
	client.htlcInterceptorRequests <- interceptReq
	resp = <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_RESUME, resp.Action)

	// Third htlc again right after should be rejected.
	interceptReq.IncomingCircuitKey.HtlcId++
	client.htlcInterceptorRequests <- interceptReq
	resp = <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_FAIL, resp.Action)

	htlcEvent := &routerrpc.HtlcEvent{
		EventType:         routerrpc.HtlcEvent_FORWARD,
		IncomingChannelId: key.ChanId,
		IncomingHtlcId:    key.HtlcId,
		Event:             &routerrpc.HtlcEvent_ForwardFailEvent{},
	}

	client.htlcEvents <- htlcEvent

	// Allow some time for the peer controller to process the failed forward
	// event.
	time.Sleep(time.Second)

	cancel()
	require.ErrorIs(t, <-exit, context.Canceled)
}

func TestMaxPending(t *testing.T) {
	defer Timeout()()

	cfg := &config{
		groupConfig: groupConfig{
			HtlcMinInterval: time.Minute,
			HtlcBurstSize:   2,
			MaxPendingHtlcs: 1,
		},
	}

	client := newLndclientMock()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newProcess(client, cfg)

	exit := make(chan error)
	go func() {
		exit <- p.run(ctx)
	}()

	key := &routerrpc.CircuitKey{
		ChanId: 2,
		HtlcId: 5,
	}
	interceptReq := &routerrpc.ForwardHtlcInterceptRequest{
		IncomingCircuitKey: key,
	}

	// First htlc accepted.
	client.htlcInterceptorRequests <- interceptReq
	resp := <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_RESUME, resp.Action)

	// Second htlc should be failed because of the max pending htlcs limit.
	interceptReq.IncomingCircuitKey.HtlcId++
	client.htlcInterceptorRequests <- interceptReq

	resp = <-client.htlcInterceptorResponses
	require.Equal(t, routerrpc.ResolveHoldForwardAction_FAIL, resp.Action)

	cancel()
	require.ErrorIs(t, <-exit, context.Canceled)
}
