package main

import "os"

func main() {
	if err := cli.Run(os.Args); err != nil {
		os.Exit(0)
	}
}
