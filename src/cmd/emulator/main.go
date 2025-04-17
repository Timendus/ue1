package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/timendus/ue1/helpers"
	"github.com/timendus/ue1/ue1"
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

	var program []byte
	switch path.Ext(filename) {
	case ".asm":
		// Assemble the source code
		fmt.Printf("Emulating file '%s'\n", filename)
		var err error
		contents, err := helpers.LoadTextFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		program, err = ue1.Assemble(contents)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case ".bin":
		// Load the binary file to execute
		fmt.Printf("Emulating file '%s'\n", filename)
		var err error
		program, err = helpers.LoadBinaryFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("\033[91;1mUnknown file type\033[0m: " + os.Args[1] + "\nExpected a .asm or a .bin file.")
		os.Exit(1)
	}

	// Actually run the emulator
	run(&ue1.UE1{
		Program:  program,
		Speed:    speed,
		State:    ue1.STATE_RUNNING,
		BellFunc: func() { fmt.Print("\a") },
	})
}
