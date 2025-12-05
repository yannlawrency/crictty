package main

import (
	"os"

	"github.com/yannlawrency/crictty/cmd"
)

var version = "dev"

// main is the entry point of the application
func main() {
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
