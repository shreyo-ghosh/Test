package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestHandleDeploy(t *testing.T) {
	tests := []struct {
		name        string
		cmd         deployCmd
		shouldError bool
	}{
		{
			name: "Valid deployment",
			cmd: deployCmd{
				function:    "test-function",
				environment: "dev",
				revision:    "1.0.0",
				clean:       true,
			},
			shouldError: false,
		},
		{
			name: "Missing environment",
			cmd: deployCmd{
				function: "test-function",
			},
			shouldError: true,
		},
		{
			name: "Invalid environment",
			cmd: deployCmd{
				function:    "test-function",
				environment: "invalid",
			},
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
				handleDeploy(tt.cmd)
				exitChan <- 0
			}()

			// Get the exit code
			exitCode := <-exitChan

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Check if the test should have errored
			if tt.shouldError && exitCode == 0 {
				t.Errorf("handleDeploy() should have errored but didn't")
			} else if !tt.shouldError && exitCode != 0 {
				t.Errorf("handleDeploy() errored when it shouldn't have")
			}
		})
	}
}

func TestHandleDescribe(t *testing.T) {
	tests := []struct {
		name string
		cmd  describeCmd
	}{
		{
			name: "Valid function",
			cmd: describeCmd{
				function: "test-function",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			handleDescribe(tt.cmd)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read the output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Check if the output contains the function name
			if !strings.Contains(output, tt.cmd.function) {
				t.Errorf("handleDescribe() output does not contain function name: %s", output)
			}
		})
	}
}

func TestPrintHelp(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printHelp()

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
			t.Errorf("printHelp() output does not contain section: %s", section)
		}
	}
}

// Mock os.Exit for testing
var osExit = os.Exit 