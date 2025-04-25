package main

import (
	"flag"
	"fmt"
	"os"
)

type deployCmd struct {
	function    string
	environment string
	revision    string
	clean       bool
}

type describeCmd struct {
	function string
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		deploy := deployCmd{}
		fs := flag.NewFlagSet("deploy", flag.ExitOnError)
		fs.StringVar(&deploy.environment, "e", "", "Target environment (sandbox, dev, or pro)")
		fs.StringVar(&deploy.revision, "v", "", "Revision number")
		fs.BoolVar(&deploy.clean, "c", false, "Clean and rebuild before deploying")
		fs.Parse(os.Args[2:])

		if fs.NArg() != 1 {
			fmt.Println("Error: Function name is required")
			printHelp()
			os.Exit(1)
		}
		deploy.function = fs.Arg(0)
		handleDeploy(deploy)

	case "describe":
		describe := describeCmd{}
		fs := flag.NewFlagSet("describe", flag.ExitOnError)
		fs.Parse(os.Args[2:])

		if fs.NArg() != 1 {
			fmt.Println("Error: Function name is required")
			printHelp()
			os.Exit(1)
		}
		describe.function = fs.Arg(0)
		handleDescribe(describe)

	case "-h", "--help", "help":
		printHelp()
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func handleDeploy(cmd deployCmd) {
	if cmd.environment == "" {
		fmt.Println("Error: Environment (-e) is required")
		printHelp()
		os.Exit(1)
	}

	if cmd.environment != "sandbox" && cmd.environment != "dev" && cmd.environment != "pro" {
		fmt.Println("Error: Environment must be one of: sandbox, dev, pro")
		os.Exit(1)
	}

	fmt.Printf("Deploying function '%s' to %s environment\n", cmd.function, cmd.environment)
	if cmd.revision != "" {
		fmt.Printf("Using revision: %s\n", cmd.revision)
	}
	if cmd.clean {
		fmt.Println("Cleaning and rebuilding before deployment...")
	}
	// TODO: Implement actual deployment logic
}

func handleDescribe(cmd describeCmd) {
	fmt.Printf("Getting details for function '%s'\n", cmd.function)
	// TODO: Implement actual description logic
}

func printHelp() {
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