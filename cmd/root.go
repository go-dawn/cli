package cmd

import (
	"github.com/spf13/cobra"
)

const version = "v0.0.7"

func init() {
	rootCmd.AddCommand(
		versionCmd, newCmd, generateCmd, devCmd,
	)
}

var rootCmd = &cobra.Command{
	Use:   "dawn",
	Short: "Dawn is an opinionated lightweight framework",
	Long:  longDescription,
	RunE:  rootRunE,
}

func rootRunE(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func Run() (err error) {
	return rootCmd.Execute()
}

const longDescription = `       __
   ___/ /__ __    _____    dawn-cli ` + version + `
 ~/ _  / _ '/ |/|/ / _ \~  For the opinionated lightweight framework dawn
~~\_,_/\_,_/|__,__/_//_/~~ Visit https://github.com/go-dawn/dawn for detail
 ~~~  ~~ ~~~~~~~~~ ~~~~~~  (c) since 2020 by kiyonlin@gmail.com`
