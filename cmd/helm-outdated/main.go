package main

import (
	"fmt"
	"os"
)

func main() {
	if err := NewCommand().Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
