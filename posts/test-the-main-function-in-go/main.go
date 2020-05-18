package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// I'm ok with not testing this call
	os.Exit(realMain(os.Stdout))
}

func realMain(out io.Writer) int {
	name := flag.String("name", "", "Your Name")
	flag.Parse()

	if *name == "" {
		fmt.Fprintf(out, "Missing flag -name\n")
		return 1
	}

	fmt.Fprintf(out, "Hi %v\n", *name)
	return 0
}
