package main

import (
	"os"

	"github.com/lstyles/nsgflowlogsbeat/cmd"

	_ "github.com/lstyles/nsgflowlogsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
