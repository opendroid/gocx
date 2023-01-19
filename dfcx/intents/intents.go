package intents

import (
	"fmt"

	"cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
)

// ShowIntents prints the intents to stdout.
func ShowIntents(intents []*cxpb.Intent, verbose bool) {
	for _, intent := range intents {
		if verbose {
			fmt.Printf("ShowIntents: %s\n", intent.GetDisplayName())
			fmt.Printf("ShowIntents: Training phrases:\n")
			for _, phrase := range intent.GetTrainingPhrases() {
				for _, part := range phrase.GetParts() {
					fmt.Printf("  %s\n", part.GetText())
				}
			}
		} else {
			fmt.Printf("%s\n", intent.DisplayName)
		}
	}
}
