package main

import (
	"fmt"
	"os"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	if err := NewCommand().Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
