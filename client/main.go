package main

import (
	"fmt"
	"os"
)

func main() {

	conf, err := setupParseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errors := validateConf(conf)
	if len(errors) != 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}
	output, err := performAction(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, output)

}
