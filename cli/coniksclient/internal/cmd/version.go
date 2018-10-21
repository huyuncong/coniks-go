package cmd

import (
	"github.com/huyuncong/coniks-go/cli"
)

var versionCmd = cli.NewVersionCommand("coniksclient")

func init() {
	RootCmd.AddCommand(versionCmd)
}
