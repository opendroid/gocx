package clients

import (
	"context"
	"os"
	"time"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"github.com/opendroid/gocx/log"
	"google.golang.org/api/iterator"
)

// Environment is Dialogflow CX bot environments
type Environment string

// Dialog flow environmental names
const (
	Draft   Environment = "DRAFT" // DRAFT environment
	Staging Environment = "STAGING"
	Prod    Environment = "PROD"
)

// String returns the string representation of the environment
func (e Environment) String() string {
	return string(e)
}

const (
	// GCPProjectID  GCP project ID in use
	GCPProjectID = "your-gcp-project-id"
	// DefaultLanguage of the bot
	DefaultLanguage = "en"
	// ENUSLanguage US English
	ENUSLanguage = "en-US"
	// DefaultTimeZone where user is in
	DefaultTimeZone = "PST"
	// DefaultTimeout for API timeout
	DefaultTimeout = 10 * time.Second
)

var (
	// _credentialsFile is the path to the JSON credentials file.
	_credentialsFile string
)

func init() {
	_credentialsFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if _credentialsFile == "" {
		log.Sugar.Fatal("GOOGLE_APPLICATION_CREDENTIALS not set")
	}
}

// cxAgent is the Dialogflow CX agent.
type cxAgent struct {
	intentsClient *cx.IntentsClient
	flowClient    *cx.FlowsClient
	pageClient    *cx.PagesClient
	agentClient   *cx.AgentsClient
}

// CXAgentOpts options passed to the agent.
type CXAgentOpts struct {
	ProjectID   string
	AgentID     string
	Verbose     bool
	Lang        string
	Location    string
	OutFilename string
	InFilename  string
}

func New() CX {
	agent := &cxAgent{
		intentsClient: NewIntents(),
		flowClient:    NewFlows(),
		pageClient:    NewPages(),
		agentClient:   NewAgents(),
	}
	return agent
}

// Implement the CX interface

// GetIntents returns a list of intents.
func (a *cxAgent) GetIntents(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Intent, error) {
	if a == nil {
		return nil, nil
	}
	// Reference: https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3/cxpb#ListIntentsRequest
	// projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>
	parent := "projects/" + opts.ProjectID + "/locations/" + opts.Location + "/agents/" + opts.AgentID
	req := &cxpb.ListIntentsRequest{
		Parent:       parent,
		PageSize:     1000,
		LanguageCode: opts.Lang,
	}
	it := a.intentsClient.ListIntents(ctx, req)
	var intents []*cxpb.Intent
	for {
		intent, err := it.Next()
		if err == iterator.Done {
			return intents, nil
		} else if err != nil {
			return nil, err
		}
		if opts.Verbose {
			log.Sugar.Infow("GetIntents", "intent", intent.DisplayName, "training phrases", len(intent.TrainingPhrases), "parameters", len(intent.Parameters))
		}
		intents = append(intents, intent)
	}
}

// GetFlows returns a list of flows.
func (a *cxAgent) GetFlows(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Flow, error) {
	if a == nil {
		return nil, nil
	}
	if opts.Verbose {
		log.Sugar.Infow("GetFlows", "agent", opts.AgentID)
	}
	return nil, nil
}

// GetPages returns a list of pages.
func (a *cxAgent) GetPages(ctx context.Context, opts *CXAgentOpts) ([]*cxpb.Page, error) {
	if a == nil {
		return nil, nil
	}
	return nil, nil
}

// GetAgent returns the agent.
func (a *cxAgent) GetAgent(ctx context.Context, opts *CXAgentOpts) (*cxpb.Agent, error) {
	if a == nil {
		return nil, nil
	}
	return nil, nil
}

// Close closes the agent.
func (a *cxAgent) Close() {
	a.intentsClient.Close()
	a.flowClient.Close()
	a.pageClient.Close()
	a.agentClient.Close()
}
