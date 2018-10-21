package cmd

import (
	"github.com/huyuncong/coniks-go/cli"
)

var versionCmd = cli.NewVersionCommand("coniksserver")

func init() {
	RootCmd.AddCommand(versionCmd)
}
