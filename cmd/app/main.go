package main

import (
	gotenbergcmd "github.com/gotenberg/gotenberg/v8/cmd"
	_ "github.com/gotenberg/gotenberg/v8/pkg/standard"

	// TODO: change namespace.
	_ "github.com/diekruecke/gotenberg-paperless/pkg/modules/example"
)

func main() {
	gotenbergcmd.Run()
}
