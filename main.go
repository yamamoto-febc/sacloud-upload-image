package main

import (
	"github.com/yamamoto-febc/sacloud-upload-image/cli"
	"os"
)

func main() {
	cli := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
