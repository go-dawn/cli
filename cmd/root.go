package cmd

import (
	"strings"

	"github.com/muesli/termenv"
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
	Long:  colorizeLongDescription(),
	RunE:  rootRunE,
}

func rootRunE(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func Run() (err error) {
	return rootCmd.Execute()
}

func colorizeLongDescription() string {
	var oldnew []string

	yellow := `_'/\|,`
	for _, c := range yellow {
		logo := termenv.String(string(c)).
			Foreground(termenv.ANSIYellow).String()
		oldnew = append(oldnew, string(c), logo)
	}

	wave := termenv.String("~").
		Foreground(termenv.ANSICyan).String()
	oldnew = append(oldnew, "~", wave)

	v := termenv.String(version).
		Foreground(termenv.ANSIYellow).String()
	oldnew = append(oldnew, version, v)

	github := termenv.String("https://github.com/go-dawn/dawn").
		Foreground(termenv.ANSICyan).
		Underline().Italic().String()
	oldnew = append(oldnew, "https://github.com/go-dawn/dawn", github)

	replacer := strings.NewReplacer(oldnew...)

	return replacer.Replace(longDescription)
}

const longDescription = `       __
   ___/ /__ __    _____    dawn-cli ` + version + `
 ~/ _  / _ '/ |/|/ / _ \~  For the opinionated lightweight framework dawn
~~\_,_/\_,_/|__,__/_//_/~~ Visit https://github.com/go-dawn/dawn for detail
 ~~~  ~~ ~~~~~~~~~ ~~~~~~  (c) since 2020 by kiyonlin@gmail.com`
