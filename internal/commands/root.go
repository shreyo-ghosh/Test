package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gcptool",
		Short: "GCP Cloud Function deployment tool",
		Long: `A command line tool for deploying and managing GCP Cloud Functions.
Complete documentation is available at https://github.com/carbonquest/gcptool`,
	}

	cmd.AddCommand(NewDeployCommand())
	cmd.AddCommand(NewDescribeCommand())

	return cmd
} 