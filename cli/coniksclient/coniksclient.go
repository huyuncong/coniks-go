// Executable CONIKS test client. See README for
// usage instructions.
package main

import (
	"github.com/huyuncong/coniks-go/cli"
	"github.com/huyuncong/coniks-go/cli/coniksclient/internal/cmd"
)

func main() {
	cli.Execute(cmd.RootCmd)
}
