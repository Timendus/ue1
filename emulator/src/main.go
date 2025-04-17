package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Parse command line parameters
	if len(os.Args) < 2 {
		fmt.Println("\033[91;1mInput file is a required parameter\033[0m\nUsage:\n   ue1emu <input file> [<cpu speed in hz>]")
		os.Exit(1)
	}
	filename := os.Args[1]
	var speed int
	if len(os.Args) > 2 {
		var err error
		speed, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("\033[91;1mInvalid cpu speed\033[0m: " + os.Args[2] + " (should be integer)")
			os.Exit(1)
		}
	} else {
		speed = 50
	}

	// Load the binary file to execute
	fmt.Printf("Emulating file '%s'\n", filename)
	program, err := loadFile(filename)
	if err != nil {
		panic(err)
	}

	// Actually run the emulator
	cpu := UE1{
		program: program,
		speed:   speed,
		running: true,
		state:   STATE_RUNNING,
	}
	cpu.run()
}

func loadFile(filename string) ([]byte, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("Requested file '%s' not found", filename)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file '%s': %s", filename, err.Error())
	}
	return file, nil
}
