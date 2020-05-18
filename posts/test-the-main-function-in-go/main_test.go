package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestFlags(T *testing.T) {
	// We manipuate the Args to set them up for the testcases
	// After this test we restore the initial args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		Name           string
		Args           []string
		ExpectedExit   int
		ExpectedOutput string
	}{
		{"flags set", []string{"-name", "johannes"}, 0, "Hi johannes\n"},
		{"flags not set", []string{"test"}, 1, "Missing flag -name\n"},
	}

	for _, tc := range cases {
		// this call is required because otherwise flags panics,
		// if args are set between flag.Parse call
		flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		os.Args = append([]string{tc.Name}, tc.Args...)
		var buf bytes.Buffer
		actualExit := realMain(&buf)
		if tc.ExpectedExit != actualExit {
			T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}

		actualOutput := buf.String()
		if tc.ExpectedOutput != actualOutput {
			T.Errorf("Wrong output for args: %v, expected %v, got: %v",
				tc.Args, tc.ExpectedOutput, actualOutput)
		}

	}

}
