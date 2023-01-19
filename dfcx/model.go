package dfcx

import (
	"context"

	"github.com/opendroid/gocx/clients"
	"github.com/opendroid/gocx/dfcx/intents"
	"github.com/opendroid/gocx/log"
)

// Model is the interface for the Dialogflow CX model. Used by cmd package.
type Model interface {
	ShowIntents(projectID, locationID, agentID, languageCode string, verbose bool)
	SaveIntents(projectID, locationID, agentID, languageCode, filename string, verbose bool)
}

// dfcx is the struct that satisfies Model interface.
type dfcx struct {
	agent clients.CX
}

func New() Model {
	dfcx := &dfcx{
		agent: clients.New(),
	}
	return dfcx
}

// ShowIntents prints the intents to stdout.
func (dfcx *dfcx) ShowIntents(projectID, locationID, agentID, languageCode string, verbose bool) {
	ctx, cancel := context.WithTimeout(context.Background(), clients.DefaultTimeout)
	defer cancel()
	opts := &clients.CXAgentOpts{
		ProjectID: projectID,
		Location:  locationID,
		AgentID:   agentID,
		Lang:      languageCode,
		Verbose:   verbose,
	}
	if _intents, err := dfcx.agent.GetIntents(ctx, opts); err != nil {
		log.Sugar.Errorw("ShowIntents", "error", err.Error())
	} else {
		intents.ShowIntents(_intents, verbose)
	}
}

// SaveIntents saves the intents to a file.
func (dfcx *dfcx) SaveIntents(projectID, locationID, agentID, languageCode, filename string, verbose bool) {
}
