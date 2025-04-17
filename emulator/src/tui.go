package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

func (cpu *UE1) run() {
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
				cpu.running = false
				return
			}

			switch char {
			case 's':
				cpu.state = STATE_STEP
			case 'r':
				cpu.state = STATE_RUNNING
			case 'h':
				cpu.state = STATE_HALTED
			case '1':
				cpu.ir = cpu.ir ^ (1 << 7)
			case '2':
				cpu.ir = cpu.ir ^ (1 << 6)
			case '3':
				cpu.ir = cpu.ir ^ (1 << 5)
			case '4':
				cpu.ir = cpu.ir ^ (1 << 4)
			case '5':
				cpu.ir = cpu.ir ^ (1 << 3)
			case '6':
				cpu.ir = cpu.ir ^ (1 << 2)
			case '7':
				cpu.ir = cpu.ir ^ (1 << 1)
			}
		}
	}()

	// Run the emulator until we exit the program
	for cpu.running {
		cpu.printInternals()

		if cpu.state != STATE_HALTED {
			cpu.step()

			if cpu.state == STATE_STEP {
				// Make sure we halt again after each step
				cpu.state = STATE_HALTED
			} else {
				// Otherwise, throttle CPU speed
				time.Sleep(time.Duration(1000/cpu.speed) * time.Millisecond)
			}
		} else {
			// Don't hog the host CPU if we have nothing to do
			time.Sleep(time.Duration(50) * time.Millisecond)
		}
	}
}

func (cpu *UE1) printInternals() {
	fmt.Print("\033[2J\033[H") // Clear the screen

	fmt.Printf("Next Instruction : %08b\n", cpu.program[cpu.pc]&0xF0)
	fmt.Println("Memory address   : " + strconv.Itoa(cpu.pc))
	fmt.Println("State            : " + cpu.stateString())
	fmt.Println()

	fmt.Println("== REGISTERS ==")
	fmt.Print("Input enable     = ")
	if cpu.ien {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Output enable    = ")
	if cpu.oen {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Carry            = ")
	if cpu.carry {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("Result register  = ")
	if cpu.rr {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
	fmt.Printf("Scratch register = %08b\n", cpu.sr)
	fmt.Printf("Output register  = %08b\n", cpu.or)
	fmt.Printf("Input switches   = %08b\n", cpu.ir)
	fmt.Println()
	fmt.Println("Keys: [H]alt [S]tep [R]un. [1]-[7] to toggle input switches. [Escape] to quit.")
}

func (cpu *UE1) stateString() string {
	switch cpu.state {
	case STATE_HALTED:
		return "HALTED"
	case STATE_RUNNING:
		return "RUNNING"
	case STATE_STEP:
		return "STEPPING"
	default:
		fail("These should be the only states")
	}
	return ""
}
