package gcp

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

type CloudFunctionClient struct {
	service *cloudfunctions.Service
	project string
	region  string
}

type FunctionDetails struct {
	Name         string
	Status       string
	Version      string
	LastModified time.Time
	Runtime      string
	Environment  string
}

func NewCloudFunctionClient() (*CloudFunctionClient, error) {
	ctx := context.Background()

	// First, get the current project and region from gcloud
	project, err := getGCloudConfig("project")
	if err != nil {
		return nil, fmt.Errorf("failed to get GCP project: %v", err)
	}

	region, err := getGCloudConfig("region")
	if err != nil {
		return nil, fmt.Errorf("failed to get GCP region: %v", err)
	}

	// Initialize the Cloud Functions service
	service, err := cloudfunctions.NewService(ctx, option.WithScopes(cloudfunctions.CloudPlatformScope))
	if err != nil {
		return nil, fmt.Errorf("failed to create Cloud Functions service: %v", err)
	}

	return &CloudFunctionClient{
		service: service,
		project: project,
		region:  region,
	}, nil
}

func (c *CloudFunctionClient) DeployFunction(functionName, environment, revision string, clean bool) error {
	// Validate environment
	if !isValidEnvironment(environment) {
		return fmt.Errorf("invalid environment: %s", environment)
	}

	// If clean build is requested, delete existing function first
	if clean {
		if err := c.deleteFunction(functionName); err != nil {
			return fmt.Errorf("failed to delete existing function: %v", err)
		}
	}

	// Build the function
	if err := c.buildFunction(functionName); err != nil {
		return fmt.Errorf("failed to build function: %v", err)
	}

	// Deploy the function
	cmd := exec.Command("gcloud", "functions", "deploy", functionName,
		"--runtime", "go121",
		"--trigger-http",
		"--allow-unauthenticated",
		"--project", c.project,
		"--region", c.region,
		"--env-vars-file", fmt.Sprintf("env.%s.yaml", environment))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to deploy function: %v\nOutput: %s", err, string(output))
	}

	log.Printf("Successfully deployed function %s", functionName)
	return nil
}

func (c *CloudFunctionClient) DescribeFunction(functionName string) (*FunctionDetails, error) {
	// Get function details using gcloud
	cmd := exec.Command("gcloud", "functions", "describe", functionName,
		"--project", c.project,
		"--region", c.region,
		"--format", "json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to describe function: %v\nOutput: %s", err, string(output))
	}

	// Parse the output to get function details
	// This is a simplified version - in a real implementation, you'd parse the JSON output
	details := &FunctionDetails{
		Name:         functionName,
		Status:       "ACTIVE", // This would come from the actual output
		Version:      "1.0.0",  // This would come from the actual output
		LastModified: time.Now(),
		Runtime:      "go121",
		Environment:  "dev", // This would come from the actual output
	}

	return details, nil
}

// Helper functions
func getGCloudConfig(configName string) (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", configName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s: %v", configName, err)
	}
	return strings.TrimSpace(string(output)), nil
}

func (c *CloudFunctionClient) deleteFunction(functionName string) error {
	cmd := exec.Command("gcloud", "functions", "delete", functionName,
		"--project", c.project,
		"--region", c.region,
		"--quiet")

	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to delete function: %v", err)
	}
	return nil
}

func (c *CloudFunctionClient) buildFunction(functionName string) error {
	// Build the Go function
	cmd := exec.Command("go", "build", "-o", "function", "./cmd/function")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build function: %v\nOutput: %s", err, string(output))
	}
	return nil
}

func isValidEnvironment(env string) bool {
	validEnvs := map[string]bool{
		"sandbox": true,
		"dev":     true,
		"pro":     true,
	}
	return validEnvs[env]
}
