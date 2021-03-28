package main

import (
	"github.com/FixIT-hackathon/meta-transfer-from/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(0)
	}
}
