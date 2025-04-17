package main

import (
	"fmt"
	"os"
	"time"
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

	contents, err := loadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Assembling '%s' 🡆 '%s'\n", input, output)

	result, err := assemble(contents)
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

func loadFile(filename string) (string, error) {
	if _, err := os.Stat(filename); err != nil {
		return "", fmt.Errorf("Requested file '%s' not found", filename)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Error reading file '%s': %s", filename, err.Error())
	}
	return string(file), nil
}
