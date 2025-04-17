package main

import (
	"fmt"
	"os"
	"time"

	"github.com/timendus/ue1/helpers"
	"github.com/timendus/ue1/ue1"
)

func main() {
	startTime := time.Now()

	// Parse parameters
	if len(os.Args) < 3 {
		fmt.Println("\033[91;1mInput and output file are required parameters\033[0m\nUsage:\n   ue1asm <input file> <ouput file>")
		os.Exit(1)
	}
	input := os.Args[1]
	output := os.Args[2]

	contents, err := helpers.LoadTextFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Assembling '%s' 🡆 '%s'\n", input, output)

	result, err := ue1.Assemble(contents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(output, result, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\033[92;1mFinished assembling in %s\033[0m\n", time.Since(startTime))
}
