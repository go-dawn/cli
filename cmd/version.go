package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/spf13/cobra"
)

// versionCmd prints the version number of dawn
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of dawn",
	Long:  `Print the local and released version number of dawn`,
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, _ []string) {
	var (
		cur, latest string
		err         error
		w           = cmd.OutOrStdout()
	)

	if cur, err = currentVersion(); err != nil {
		cur = err.Error()
	}

	if latest, err = latestVersion(); err != nil {
		_, _ = fmt.Fprintf(w, "dawn version: %v\n", err)
		return
	}

	_, _ = fmt.Fprintf(w, "dawn version: %s(latest %s)\n", cur, latest)
}

var currentVersionRegexp = regexp.MustCompile(`github\.com/go-dawn/dawn*?\s+(.*)\n`)
var currentVersionFile = "go.mod"

func currentVersion() (string, error) {
	b, err := ioutil.ReadFile(currentVersionFile)
	if err != nil {
		return "", err
	}

	if submatch := currentVersionRegexp.FindSubmatch(b); len(submatch) == 2 {
		return string(submatch[1]), nil
	}

	return "", errors.New("github.com/go-dawn/dawn was not found in go.mod")
}

var latestVersionRegexp = regexp.MustCompile(`"name":"(v.*?)"`)

func latestVersion() (v string, err error) {
	var (
		res *http.Response
		b   []byte
	)

	if res, err = http.Get("https://api.github.com/repos/go-dawn/dawn/releases/latest"); err != nil {
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	if submatch := latestVersionRegexp.FindSubmatch(b); len(submatch) == 2 {
		return string(submatch[1]), nil
	}

	return "", errors.New("no version found in github response body")
}
