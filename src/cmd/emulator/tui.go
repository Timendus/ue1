package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/timendus/ue1/helpers"
	"github.com/timendus/ue1/ue1"
)

func run(cpu *ue1.UE1) {
	running := true

	// Make sure we read key strokes during execution
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			switch key {
			case keyboard.KeyEsc:
				running = false
				return
			}

			switch char {
			case 's':
				cpu.State = ue1.STATE_STEP
			case 'r':
				cpu.State = ue1.STATE_RUNNING
			case 'h':
				cpu.State = ue1.STATE_HALTED
			case '1':
				cpu.IR = cpu.IR ^ (1 << 7)
			case '2':
				cpu.IR = cpu.IR ^ (1 << 6)
			case '3':
				cpu.IR = cpu.IR ^ (1 << 5)
			case '4':
				cpu.IR = cpu.IR ^ (1 << 4)
			case '5':
				cpu.IR = cpu.IR ^ (1 << 3)
			case '6':
				cpu.IR = cpu.IR ^ (1 << 2)
			case '7':
				cpu.IR = cpu.IR ^ (1 << 1)
			}
		}
	}()

	// Run the emulator until we exit the program
	for running {
		printInternals(cpu)

		if cpu.State != ue1.STATE_HALTED {
			cpu.Step()

			if cpu.State == ue1.STATE_STEP {
				// Make sure we halt again after each step
				cpu.State = ue1.STATE_HALTED
			} else {
				// Otherwise, throttle CPU speed
				time.Sleep(time.Duration(1000/cpu.Speed) * time.Millisecond)
			}
		} else {
			// Don't hog the host CPU if we have nothing to do
			time.Sleep(time.Duration(50) * time.Millisecond)
		}
	}
}

func printInternals(cpu *ue1.UE1) {
	fmt.Print("\033[2J\033[H") // Clear the screen

	fmt.Printf("Next Instruction : %08b (%s)\n", cpu.Program[cpu.PC], ue1.DisassembleInstruction(cpu.Program[cpu.PC]))
	fmt.Println("Memory address   : " + strconv.Itoa(cpu.PC))
	fmt.Println("State            : " + stateString(cpu))
	fmt.Println()

	fmt.Println("== REGISTERS ==")
	fmt.Print("Input enable     = ")
	if cpu.IEN {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Output enable    = ")
	if cpu.OEN {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Carry            = ")
	if cpu.Carry {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Result register  = ")
	if cpu.RR {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
	fmt.Printf("Scratch register = %08b (%d)\n", cpu.SR, cpu.SR)
	fmt.Printf("Output register  = %08b (%d)\n", cpu.OR, cpu.OR)
	fmt.Printf("Input switches   = %07b (%d)\n", cpu.IR>>1, cpu.IR>>1)
	fmt.Println()
	fmt.Println("Keys: [H]alt [S]tep [R]un. [1]-[7] to toggle input switches. [Escape] to quit.")
}

func stateString(cpu *ue1.UE1) string {
	switch cpu.State {
	case ue1.STATE_HALTED:
		return "HALTED"
	case ue1.STATE_RUNNING:
		return "RUNNING"
	case ue1.STATE_STEP:
		return "STEPPING"
	default:
		helpers.Fail("These should be the only states")
	}
	return ""
}
