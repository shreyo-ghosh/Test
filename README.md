# GCP Tool

A command-line tool for managing Google Cloud Functions, written in Go.

## Prerequisites

- Go 1.21 or later
- Google Cloud SDK installed and configured

## Installation

1. Clone this repository
2. Build the binary:
   ```bash
   go build -o gcptool
   ```
3. (Optional) Move the binary to a directory in your PATH:
   ```bash
   mv gcptool /usr/local/bin/
   ```

## Usage

### Deploy a Function

```bash
gcptool deploy <function> -e <environment> [-v <revision>] [-c]
```

Options:
- `-e`: Target environment (sandbox, dev, or pro)
- `-v`: Optional revision number
- `-c`: Clean and rebuild before deploying

Examples:
```bash
gcptool deploy my-function -e dev -v 1.0.0
gcptool deploy my-function -e pro -c
```

### Describe a Function

```bash
gcptool describe <function>
```

Example:
```bash
gcptool describe my-function
```

### Help

To see the help message:
```bash
gcptool -h
```

## Development

To add new dependencies:
```bash
go get <package>
```

To update dependencies:
```bash
go mod tidy
``` 