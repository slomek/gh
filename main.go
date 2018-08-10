package main

import (
	"fmt"
	"os"

	"github.com/slomek/gh/cmd"
)

func main() {
	if err := cmd.Command.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
}
