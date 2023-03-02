package clients

import (
	"context"

	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
)

// CX is the interface for accessing the Dialogflow CX client.
// Generate the mock interface for this.
// mockgen -destination=mocks/clients/mock_cx.go -package=mocks -source=clients/get.go CX
type CX interface {
	GetIntents(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Intent, error)
	GetFlows(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Flow, error)
	GetPages(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Page, error)
	GetAgent(ctx context.Context, opts *CXAgentOpts) (*cxpb.Agent, error)
	Close()
}
