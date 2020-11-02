package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-dawn/cli/cmd/internal"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

const version = "0.0.7"
const configName = ".dawnconfig"

var (
	rc = rootConfig{
		CliVersionCheckInterval: int64((time.Hour * 12) / time.Second),
	}

	verbose bool
)

type rootConfig struct {
	CliVersionCheckInterval int64 `json:"cli_version_check_interval"`
	CliVersionCheckedAt     int64 `json:"cli_version_checked_at"`
}

func init() {
	rootCmd.AddCommand(
		versionCmd, newCmd, generateCmd, devCmd, upgradeCmd,
	)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:               "dawn",
	Short:             "Dawn is an opinionated lightweight framework",
	Long:              colorizeLongDescription(),
	RunE:              rootRunE,
	PersistentPreRun:  rootPersistentPreRun,
	PersistentPostRun: rootPersistentPostRun,
	SilenceErrors:     true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_ = rootCmd.Help()
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n", err)
	}
}

func rootRunE(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func rootPersistentPreRun(cmd *cobra.Command, _ []string) {
	if err := loadConfig(); err != nil {
		warning := fmt.Sprintf("WARNING: failed to load config: %s\n\n", err)
		cmd.Println(termenv.String(warning).Foreground(termenv.ANSIBrightYellow))
	}
}

func rootPersistentPostRun(cmd *cobra.Command, _ []string) {
	checkCliVersion(cmd)
}

func checkCliVersion(cmd *cobra.Command) {
	if !needCheckCliVersion() {
		return
	}

	cliLatestVersion, err := latestVersion(true)
	if err != nil {
		return
	}

	if version != cliLatestVersion {
		title := termenv.String(fmt.Sprintf(versionUpgradeTitleFormat, version, cliLatestVersion)).
			Foreground(termenv.ANSIBrightYellow)

		prompt := internal.NewPrompt(title.String())
		ok, err := prompt.YesOrNo()

		if err == nil && ok {
			upgrade(cmd, cliLatestVersion)
		}

		if err != nil {
			warning := fmt.Sprintf("WARNING: Failed to upgrade dawn cli: %s", err)
			cmd.Println(termenv.String(warning).Foreground(termenv.ANSIBrightYellow))
		}
	}

	updateVersionCheckedAt()
}

func updateVersionCheckedAt() {
	rc.CliVersionCheckedAt = time.Now().Unix()
	storeConfig()
}

func needCheckCliVersion() bool {
	return !upgraded && rc.CliVersionCheckedAt+rc.CliVersionCheckInterval < time.Now().Unix()
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

	v := termenv.String("v" + version).
		Foreground(termenv.ANSIYellow).String()
	oldnew = append(oldnew, "v"+version, v)

	github := termenv.String("https://github.com/go-dawn/dawn").
		Foreground(termenv.ANSICyan).
		Underline().Italic().String()
	oldnew = append(oldnew, "https://github.com/go-dawn/dawn", github)

	replacer := strings.NewReplacer(oldnew...)

	return replacer.Replace(longDescription)
}

const (
	longDescription = `       __
   ___/ /__ __    _____    dawn-cli v` + version + `
 ~/ _  / _ '/ |/|/ / _ \~  For the opinionated lightweight framework dawn
~~\_,_/\_,_/|__,__/_//_/~~ Visit https://github.com/go-dawn/dawn for detail
 ~~~  ~~ ~~~~~~~~~ ~~~~~~  (c) since 2020 by kiyonlin@gmail.com`

	versionUpgradeTitleFormat = `
You are using dawn cli version %s; however, version %s is available.
Would you like to upgrade now? (y/N)`
)
