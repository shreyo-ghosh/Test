# GCP Cloud Function Deployment Tool

A command line tool for deploying and managing GCP Cloud Functions.

## Features

- Deploy cloud functions to different environments (sandbox, dev, pro)
- Get detailed information about deployed functions
- Support for versioning and clean builds
- GitHub Actions integration for automated deployment

## Installation

```bash
go install github.com/carbonquest/gcptool@latest
```

## Usage

### Deploy a Function

```bash
gcptool deploy <function-name> -e <environment> -v <revision> -c
```

Options:
- `-e, --environment`: Environment to deploy to (sandbox, dev, pro)
- `-v, --revision`: Revision number
- `-c, --clean`: Clean and rebuild before deploying

Example:
```bash
gcptool deploy my-function -e dev -v 1.0.0 -c
```

### Describe a Function

```bash
gcptool describe <function-name>
```

Example:
```bash
gcptool describe my-function
```

## GitHub Actions Integration

The repository includes a GitHub Actions workflow that:
1. Runs unit tests on every push and pull request
2. Deploys the function to GCP when changes are pushed to main branch

### Setup

1. Add your GCP credentials as a GitHub secret named `GCP_CREDENTIALS`
2. The workflow will automatically run on pushes to main branch

## Development

### Building from Source

```bash
git clone https://github.com/carbonquest/gcptool.git
cd gcptool
go build -o gcptool ./cmd/gcptool
```

### Running Tests

```bash
go test ./... -v
```

## License

MIT 