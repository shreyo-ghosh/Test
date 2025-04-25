package commands

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestHandleDeploy(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		shouldError bool
	}{
		{
			name:        "Valid deployment",
			args:        []string{"-e", "dev", "-v", "1.0.0", "test-function"},
			shouldError: false,
		},
		{
			name:        "Missing environment",
			args:        []string{"test-function"},
			shouldError: true,
		},
		{
			name:        "Invalid environment",
			args:        []string{"-e", "invalid", "test-function"},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Create a channel to capture the exit code
			exitChan := make(chan int)
			oldExit := osExit
			osExit = func(code int) {
				exitChan <- code
			}
			defer func() { osExit = oldExit }()

			// Run the function in a goroutine
			go func() {
				HandleDeploy(tt.args)
				exitChan <- 0
			}()

			// Get the exit code
			exitCode := <-exitChan

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Check if the test should have errored
			if tt.shouldError && exitCode == 0 {
				t.Errorf("HandleDeploy() should have errored but didn't")
			} else if !tt.shouldError && exitCode != 0 {
				t.Errorf("HandleDeploy() errored when it shouldn't have")
			}
		})
	}
}

func TestHandleDescribe(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "Valid function",
			args: []string{"test-function"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			HandleDescribe(tt.args)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read the output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Check if the output contains expected fields
			requiredFields := []string{
				"Function:",
				"Status:",
				"Version:",
				"Last Deployed:",
			}

			for _, field := range requiredFields {
				if !strings.Contains(output, field) {
					t.Errorf("HandleDescribe() output does not contain field: %s", field)
				}
			}
		})
	}
}

func TestPrintHelp(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintHelp()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check if the output contains key sections
	requiredSections := []string{
		"Usage:",
		"Commands:",
		"Examples:",
		"Deploy Options:",
	}

	for _, section := range requiredSections {
		if !strings.Contains(output, section) {
			t.Errorf("PrintHelp() output does not contain section: %s", section)
		}
	}
}

// Mock os.Exit for testing
var osExit = os.Exit 