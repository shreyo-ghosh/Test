package commands

import (
	"fmt"
	"log"

	"github.com/carbonquest/gcptool/internal/gcp"
	"github.com/spf13/cobra"
)

func NewDeployCommand() *cobra.Command {
	var (
		environment string
		revision    string
		clean       bool
	)

	cmd := &cobra.Command{
		Use:   "deploy <function>",
		Short: "Deploy a cloud function",
		Long:  `Deploy a cloud function to the specified environment with optional revision and clean build.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			functionName := args[0]
			
			// Validate environment
			if !isValidEnvironment(environment) {
				return fmt.Errorf("invalid environment: %s. Must be one of: sandbox, dev, pro", environment)
			}

			// Initialize GCP client
			client, err := gcp.NewCloudFunctionClient()
			if err != nil {
				return fmt.Errorf("failed to initialize GCP client: %v", err)
			}

			// Deploy the function
			if err := client.DeployFunction(functionName, environment, revision, clean); err != nil {
				return fmt.Errorf("failed to deploy function: %v", err)
			}

			log.Printf("Successfully deployed function %s to %s environment", functionName, environment)
			return nil
		},
	}

	cmd.Flags().StringVarP(&environment, "environment", "e", "dev", "Environment to deploy to (sandbox, dev, pro)")
	cmd.Flags().StringVarP(&revision, "revision", "v", "", "Revision number")
	cmd.Flags().BoolVarP(&clean, "clean", "c", false, "Clean and rebuild before deploying")

	return cmd
}

func isValidEnvironment(env string) bool {
	validEnvs := map[string]bool{
		"sandbox": true,
		"dev":     true,
		"pro":     true,
	}
	return validEnvs[env]
} 