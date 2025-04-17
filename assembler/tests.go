package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	testsDirectory := "../programs/"
	assembler := "../dist/linux/ue1asm"

	dir, err := os.ReadDir(testsDirectory)
	if err != nil {
		panic(err)
	}

	allGood := true
	for _, file := range dir {
		if file.IsDir() {
			inputFile := testsDirectory + file.Name() + "/" + file.Name() + ".asm"
			outputFile := testsDirectory + file.Name() + "/test-output.bin"
			expectedFile := testsDirectory + file.Name() + "/" + file.Name() + ".bin"

			if !runCommand(assembler+" "+inputFile+" "+outputFile) ||
				!runCommand("git diff --no-index --color "+expectedFile+" "+outputFile) {
				allGood = false
				fmt.Println("   ❌ " + file.Name())
			} else {
				fmt.Println("   ✔️ " + file.Name())
			}
		}
	}

	if allGood {
		fmt.Println("\033[92;1mAll good!\033[0m")
		os.Exit(0)
	} else {
		fmt.Println("\033[91;1mOne or more projects did not result in the expected output\033[0m")
		os.Exit(1)
	}
}

func runCommand(command string) bool {
	parts := strings.Split(command, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(string(output))
		return false
	}
	return true
}
