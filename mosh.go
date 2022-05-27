package main

import (
	"github.com/adamrdrew/mosh/cmd"
	"github.com/adamrdrew/mosh/config"
)

var Version string

func main() {
	config.Version = Version
	cmd.Execute()
}
