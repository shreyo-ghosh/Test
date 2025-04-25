package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yourusername/gcptool/internal/gcp"
)

type DeployOptions struct {
	Function    string
	Environment string
	Revision    string
	Clean       bool
}

type DescribeOptions struct {
	Function string
}

func HandleDeploy(args []string) {
	opts := DeployOptions{}
	fs := flag.NewFlagSet("deploy", flag.ExitOnError)
	fs.StringVar(&opts.Environment, "e", "", "Target environment (sandbox, dev, or pro)")
	fs.StringVar(&opts.Revision, "v", "", "Revision number")
	fs.BoolVar(&opts.Clean, "c", false, "Clean and rebuild before deploying")
	fs.Parse(args)

	if fs.NArg() != 1 {
		fmt.Println("Error: Function name is required")
		PrintHelp()
		os.Exit(1)
	}
	opts.Function = fs.Arg(0)

	if err := validateDeployOptions(opts); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := gcp.DeployFunction(opts); err != nil {
		fmt.Printf("Error deploying function: %v\n", err)
		os.Exit(1)
	}
}

func HandleDescribe(args []string) {
	opts := DescribeOptions{}
	fs := flag.NewFlagSet("describe", flag.ExitOnError)
	fs.Parse(args)

	if fs.NArg() != 1 {
		fmt.Println("Error: Function name is required")
		PrintHelp()
		os.Exit(1)
	}
	opts.Function = fs.Arg(0)

	details, err := gcp.DescribeFunction(opts.Function)
	if err != nil {
		fmt.Printf("Error describing function: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Function: %s\n", details.Name)
	fmt.Printf("Status: %s\n", details.Status)
	fmt.Printf("Version: %s\n", details.Version)
	fmt.Printf("Last Deployed: %s\n", details.LastDeployed)
}

func validateDeployOptions(opts DeployOptions) error {
	if opts.Environment == "" {
		return fmt.Errorf("environment (-e) is required")
	}

	if opts.Environment != "sandbox" && opts.Environment != "dev" && opts.Environment != "pro" {
		return fmt.Errorf("environment must be one of: sandbox, dev, pro")
	}

	return nil
}

func PrintHelp() {
	fmt.Println(`GCP Function Deployment Tool

Usage:
  gcptool <command> [arguments]

Commands:
  deploy    Deploy a function
  describe  Get function details
  help      Show this help message

Examples:
  gcptool deploy my-function -e dev -v 1.0.0
  gcptool deploy my-function -e pro -c
  gcptool describe my-function

Deploy Options:
  -e string  Target environment (sandbox, dev, or pro)
  -v string  Revision number
  -c         Clean and rebuild before deploying`)
} 