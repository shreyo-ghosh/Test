package gcp

import (
	"testing"
	"time"
)

func TestDescribeFunction(t *testing.T) {
	client := &CloudFunctionClient{}

	details, err := client.DescribeFunction("test-function")
	if err != nil {
		t.Fatalf("DescribeFunction failed: %v", err)
	}

	if details.Name != "test-function" {
		t.Errorf("Expected function name 'test-function', got '%s'", details.Name)
	}

	if details.Status != "ACTIVE" {
		t.Errorf("Expected status 'ACTIVE', got '%s'", details.Status)
	}

	if details.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", details.Version)
	}

	if details.Runtime != "go121" {
		t.Errorf("Expected runtime 'go121', got '%s'", details.Runtime)
	}

	if details.Environment != "dev" {
		t.Errorf("Expected environment 'dev', got '%s'", details.Environment)
	}

	// Check if LastModified is within a reasonable time range
	now := time.Now()
	if details.LastModified.After(now) || details.LastModified.Before(now.Add(-time.Hour)) {
		t.Errorf("LastModified time is not within expected range")
	}
}

func TestDeployFunction(t *testing.T) {
	client := &CloudFunctionClient{}

	// Test with clean build
	err := client.DeployFunction("test-function", "dev", "1.0.0", true)
	if err != nil {
		t.Fatalf("DeployFunction with clean build failed: %v", err)
	}

	// Test without clean build
	err = client.DeployFunction("test-function", "dev", "1.0.0", false)
	if err != nil {
		t.Fatalf("DeployFunction without clean build failed: %v", err)
	}

	// Test with invalid environment
	err = client.DeployFunction("test-function", "invalid-env", "1.0.0", false)
	if err == nil {
		t.Error("Expected error for invalid environment, got nil")
	}
} 