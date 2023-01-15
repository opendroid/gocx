package log

import (
	"fmt"
	"os"
	"testing"
)

// TestSugar tests the Sugar logger.
// go test -v -run TestSugar ./...
func TestSugar(t *testing.T) {
	// Save the current GOCX_ENV value
	env := os.Getenv("GOCX_ENV")
	for i, env := range []string{"PRODUCTION", "DEVELOPMENT"} {
		t.Run(fmt.Sprintf("TestSugar %d %s", i, env), func(t *testing.T) {
			os.Setenv("GOCX_ENV", env)
			setupSugarLogger()
			logger := Sugar
			if logger == nil {
				t.Errorf("TestSugar: Invalid logger	")
			}
		})
	}
	os.Setenv("GOCX_ENV", env) // Restore the original GOCX_ENV value
}
