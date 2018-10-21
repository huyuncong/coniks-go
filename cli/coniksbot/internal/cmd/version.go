package cmd

import (
	"github.com/huyuncong/coniks-go/cli"
)

var versionCmd = cli.NewVersionCommand("coniksbot")

func init() {
	RootCmd.AddCommand(versionCmd)
}
