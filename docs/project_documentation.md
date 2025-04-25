# CarbonQuest GCP Cloud Function Deployment Tool Documentation

## Table of Contents
1. [Project Overview](#project-overview)
2. [Architecture](#architecture)
3. [Core Components](#core-components)
4. [Implementation Details](#implementation-details)
5. [Testing Strategy](#testing-strategy)
6. [Deployment Process](#deployment-process)
7. [GitHub Actions Integration](#github-actions-integration)

## Project Overview

The CarbonQuest GCP Cloud Function Deployment Tool is a command-line utility designed to streamline the deployment and management of Google Cloud Platform (GCP) Cloud Functions. The tool provides a simple interface for developers to deploy, monitor, and manage their cloud functions across different environments.

### Key Features
- Environment-specific deployments (sandbox, dev, pro)
- Version control integration
- Automated testing and deployment
- Function status monitoring
- Clean build options

## Architecture

The project follows a modular architecture with clear separation of concerns:

```
carbonquest/
├── cmd/
│   └── gcptool/
│       └── main.go          # Entry point
├── internal/
│   ├── commands/            # CLI command implementations
│   │   ├── root.go
│   │   ├── deploy.go
│   │   └── describe.go
│   └── gcp/                 # GCP client implementation
│       ├── client.go
│       └── client_test.go
├── .github/
│   └── workflows/
│       └── deploy.yml       # GitHub Actions workflow
└── docs/
    └── project_documentation.md
```

## Core Components

### 1. Command Line Interface (CLI)

The CLI is built using the Cobra framework and provides two main commands:

#### Deploy Command
```go
gcptool deploy <function> -e <environment> -v <revision> -c
```
- `-e`: Environment (sandbox, dev, pro)
- `-v`: Version/revision number
- `-c`: Clean build flag

#### Describe Command
```go
gcptool describe <function>
```
- Retrieves current function status and details

### 2. GCP Client

The GCP client handles all interactions with Google Cloud Platform:

```go
type CloudFunctionClient struct {
    service *cloudfunctions.Service
}

type FunctionDetails struct {
    Name         string
    Status       string
    Version      string
    LastModified time.Time
    Runtime      string
    Environment  string
}
```

Key methods:
- `NewCloudFunctionClient()`: Initializes GCP connection
- `DeployFunction()`: Handles function deployment
- `DescribeFunction()`: Retrieves function details

## Implementation Details

### 1. Command Implementation

#### Root Command
```go
func NewRootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "gcptool",
        Short: "GCP Cloud Function deployment tool",
        Long:  `A command line tool for deploying and managing GCP Cloud Functions.`,
    }
    cmd.AddCommand(NewDeployCommand())
    cmd.AddCommand(NewDescribeCommand())
    return cmd
}
```

#### Deploy Command
```go
func NewDeployCommand() *cobra.Command {
    var (
        environment string
        revision    string
        clean       bool
    )
    // Command implementation
}
```

### 2. GCP Client Implementation

#### Client Initialization
```go
func NewCloudFunctionClient() (*CloudFunctionClient, error) {
    ctx := context.Background()
    service, err := cloudfunctions.NewService(ctx, 
        option.WithScopes(cloudfunctions.CloudPlatformScope))
    // Implementation details
}
```

#### Function Deployment
```go
func (c *CloudFunctionClient) DeployFunction(
    functionName, 
    environment, 
    revision string, 
    clean bool) error {
    // Deployment logic
}
```

## Testing Strategy

The project implements comprehensive testing:

### 1. Unit Tests
```go
func TestDescribeFunction(t *testing.T) {
    // Test implementation
}

func TestDeployFunction(t *testing.T) {
    // Test implementation
}
```

### 2. Integration Tests
- GitHub Actions workflow tests
- Environment validation tests
- Deployment process tests

## Deployment Process

### 1. Local Deployment
1. Build the tool
2. Configure GCP credentials
3. Run deployment command

### 2. Automated Deployment
1. Push to GitHub
2. GitHub Actions triggers
3. Tests run automatically
4. Deployment to GCP if tests pass

## GitHub Actions Integration

The project uses GitHub Actions for CI/CD:

```yaml
name: Deploy Cloud Function
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
      - name: Install dependencies
      - name: Run tests
      - name: Build gcptool
      - name: Configure GCP credentials
      - name: Deploy to GCP
```

## Best Practices

1. **Code Organization**
   - Clear package structure
   - Separation of concerns
   - Modular design

2. **Error Handling**
   - Comprehensive error checking
   - Meaningful error messages
   - Graceful failure handling

3. **Testing**
   - Unit test coverage
   - Integration testing
   - Automated testing

4. **Documentation**
   - Clear code comments
   - Comprehensive documentation
   - Usage examples

## Future Enhancements

1. **Planned Features**
   - Multiple function deployment
   - Rollback capabilities
   - Enhanced monitoring
   - Custom environment support

2. **Improvements**
   - Performance optimization
   - Enhanced error handling
   - Additional test coverage
   - Extended documentation 