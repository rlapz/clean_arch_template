package main

import (
	"fmt"
	"os"

	"github.com/rlapz/clean_arch_template/src/instance"
)

const help = `
Usage:
	go run src/main.go production
	go run src/main.go development
`

func main() {
	var isProduction bool

	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Not enough argument! \n%s\n", help)
		os.Exit(1)
	}

	if args[1] == "production" {
		isProduction = true
	} else if args[1] == "development" {
		isProduction = false
	} else {
		fmt.Fprintf(os.Stderr, "Invalid argument! \n%s\n", help)
		os.Exit(1)
	}

	if err := instance.RunApp(isProduction); err != nil {
		panic(err)
	}
}
