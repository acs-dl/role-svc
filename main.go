package main

import (
	"os"

	"github.com/acs-dl/role-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
