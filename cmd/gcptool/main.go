package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yourusername/gcptool/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		commands.PrintHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		commands.HandleDeploy(os.Args[2:])
	case "describe":
		commands.HandleDescribe(os.Args[2:])
	case "-h", "--help", "help":
		commands.PrintHelp()
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		commands.PrintHelp()
		os.Exit(1)
	}
} 