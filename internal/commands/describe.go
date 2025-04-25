package commands

import (
	"fmt"
	"log"

	"github.com/carbonquest/gcptool/internal/gcp"
	"github.com/spf13/cobra"
)

func NewDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe <function>",
		Short: "Get details about a cloud function",
		Long:  `Get detailed information about a deployed cloud function, including its current version and status.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			functionName := args[0]

			// Initialize GCP client
			client, err := gcp.NewCloudFunctionClient()
			if err != nil {
				return fmt.Errorf("failed to initialize GCP client: %v", err)
			}

			// Get function details
			details, err := client.DescribeFunction(functionName)
			if err != nil {
				return fmt.Errorf("failed to get function details: %v", err)
			}

			// Print function details
			log.Printf("Function: %s", details.Name)
			log.Printf("Status: %s", details.Status)
			log.Printf("Version: %s", details.Version)
			log.Printf("Last Modified: %s", details.LastModified)
			log.Printf("Runtime: %s", details.Runtime)
			log.Printf("Environment: %s", details.Environment)

			return nil
		},
	}

	return cmd
} 