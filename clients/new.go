package clients

import (
	"context"
	"time"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	"github.com/opendroid/gocx/log"
	"google.golang.org/api/option"
)

// NewAgentsClient returns a new AgentsClient
func NewIntents() *cx.IntentsClient {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(DefaultTimeout))
	defer cancel()
	if cx, err := cx.NewIntentsClient(ctx, option.WithCredentialsFile(_credentialsFile)); err != nil {
		log.Sugar.Fatalw("NewIntents", "error", err.Error())
		panic(err)
	} else {
		return cx
	}
}

// NewSessions returns a new SessionsClient
func NewSessions() *cx.SessionsClient {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(DefaultTimeout))
	defer cancel()
	if cx, err := cx.NewSessionsClient(ctx, option.WithCredentialsFile(_credentialsFile)); err != nil {
		log.Sugar.Fatalw("NewSessions", "error", err.Error())
		panic(err)
	} else {
		return cx
	}
}

// NewFlows returns a new FlowsClient
func NewFlows() *cx.FlowsClient {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(DefaultTimeout))
	defer cancel()
	if cx, err := cx.NewFlowsClient(ctx, option.WithCredentialsFile(_credentialsFile)); err != nil {
		log.Sugar.Fatalw("NewFlows", "error", err.Error())
		panic(err)
	} else {
		return cx
	}
}

// NewPages returns a new PagesClient
func NewPages() *cx.PagesClient {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(DefaultTimeout))
	defer cancel()
	if cx, err := cx.NewPagesClient(ctx, option.WithCredentialsFile(_credentialsFile)); err != nil {
		log.Sugar.Fatalw("NewPages", "error", err.Error())
		panic(err)
	} else {
		return cx
	}
}

// NewAgents returns a new AgentsClient
func NewAgents() *cx.AgentsClient {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(DefaultTimeout))
	defer cancel()
	if cx, err := cx.NewAgentsClient(ctx, option.WithCredentialsFile(_credentialsFile)); err != nil {
		log.Sugar.Fatalw("NewAgents", "error", err.Error())
		panic(err)
	} else {
		return cx
	}
}
