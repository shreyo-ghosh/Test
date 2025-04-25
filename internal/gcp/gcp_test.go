package gcp

import (
	"context"
	"testing"
	"time"

	"github.com/yourusername/gcptool/internal/commands"
)

func TestDeployFunction(t *testing.T) {
	tests := []struct {
		name    string
		opts    commands.DeployOptions
		wantErr bool
	}{
		{
			name: "Valid deployment",
			opts: commands.DeployOptions{
				Function:    "test-function",
				Environment: "dev",
				Revision:    "1.0.0",
				Clean:       true,
			},
			wantErr: false,
		},
		{
			name: "Invalid environment",
			opts: commands.DeployOptions{
				Function:    "test-function",
				Environment: "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeployFunction(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeployFunction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDescribeFunction(t *testing.T) {
	tests := []struct {
		name       string
		function   string
		wantErr    bool
		wantFields []string
	}{
		{
			name:     "Valid function",
			function: "test-function",
			wantErr:  false,
			wantFields: []string{
				"test-function",
				"ACTIVE",
				"1.0.0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			details, err := DescribeFunction(tt.function)
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if details.Name != tt.function {
					t.Errorf("DescribeFunction() name = %v, want %v", details.Name, tt.function)
				}
				if details.Status != "ACTIVE" {
					t.Errorf("DescribeFunction() status = %v, want %v", details.Status, "ACTIVE")
				}
				if details.Version != "1.0.0" {
					t.Errorf("DescribeFunction() version = %v, want %v", details.Version, "1.0.0")
				}
				if details.LastDeployed.IsZero() {
					t.Error("DescribeFunction() lastDeployed is zero")
				}
			}
		})
	}
}

func TestNewGCPClient(t *testing.T) {
	ctx := context.Background()
	client, err := NewGCPClient(ctx)
	if err != nil {
		t.Errorf("NewGCPClient() error = %v", err)
		return
	}
	defer client.Close()

	if client.client == nil {
		t.Error("NewGCPClient() client is nil")
	}
	if client.ctx != ctx {
		t.Error("NewGCPClient() context mismatch")
	}
} 