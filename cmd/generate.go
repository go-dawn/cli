package cmd

import "github.com/spf13/cobra"

func init() {
	generateCmd.AddCommand(
		moduleCmd,
	)
}

// generateCmd generates boilerplate code with different commands
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate boilerplate code with different commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
