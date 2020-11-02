package cmd

import (
	"fmt"
	"os"

	"github.com/go-dawn/cli/cmd/internal"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade Dawn cli if a newer version is available",
	RunE:  upgradeRunE,
}

var upgraded bool

func upgradeRunE(cmd *cobra.Command, _ []string) error {
	cliLatestVersion, err := latestVersion(true)
	if err != nil {
		return err
	}

	if version != cliLatestVersion {
		upgrade(cmd, cliLatestVersion)
	} else {
		msg := fmt.Sprintf("Currently Dawn cli is the latest version %s.", cliLatestVersion)
		cmd.Println(termenv.String(msg).
			Foreground(termenv.ANSIBrightBlue))
	}

	return nil
}

func upgrade(cmd *cobra.Command, cliLatestVersion string) {
	args := []string{"get", "-u"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, "github.com/go-dawn/cli/dawn")

	upgrader := execCommand("go", args...)
	upgrader.Env = append(upgrader.Env, os.Environ()...)
	upgrader.Env = append(upgrader.Env, "GO111MODULE=off")

	scmd := internal.NewSpinnerCmd(upgrader, "Upgrading")

	if err := scmd.Run(); err != nil && !skipSpinner {
		cmd.Printf("dawn: failed to upgrade: %v", err)
		return
	}

	success := fmt.Sprintf("Done! Dawn cli is now at v%s!", cliLatestVersion)
	cmd.Println(termenv.String(success).Foreground(termenv.ANSIBrightGreen))

	upgraded = true
}
