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

	fmt.Printf("Next Instruction : %08b (%s)\n", cpu.program[cpu.pc], opcodeToText(cpu.program[cpu.pc]))
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
	fmt.Printf("Scratch register = %08b (%d)\n", cpu.sr, cpu.sr)
	fmt.Printf("Output register  = %08b (%d)\n", cpu.or, cpu.or)
	fmt.Printf("Input switches   = %08b (%d)\n", cpu.ir, cpu.ir)
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

func opcodeToText(instruction byte) string {
	operand := instruction & 0xF
	inputOperands := []string{
		"SR0", "SR1", "SR2", "SR3", "SR4", "SR5", "SR6", "SR7", "RR", "IR1", "IR2", "IR3", "IR4", "IR5", "IR6", "IR7",
	}
	outputOperands := []string{
		"SR0", "SR1", "SR2", "SR3", "SR4", "SR5", "SR6", "SR7", "OR0", "OR1", "OR2", "OR3", "OR4", "OR5", "OR6", "OR7",
	}

	switch instruction & 0xF0 >> 4 {
	case 0b0000:
		return "NOP0"
	case 0b0001:
		return "LD " + inputOperands[operand]
	case 0b0010:
		return "ADD " + inputOperands[operand]
	case 0b0011:
		return "SUB " + inputOperands[operand]
	case 0b100:
		return "ONE"
	case 0b0101:
		return "NAND " + inputOperands[operand]
	case 0b0110:
		return "OR " + inputOperands[operand]
	case 0b0111:
		return "XOR " + inputOperands[operand]
	case 0b1000:
		return "STO " + outputOperands[operand]
	case 0b1001:
		return "STOC " + outputOperands[operand]
	case 0b1010:
		return "IEN"
	case 0b1011:
		return "OEN"
	case 0b1100:
		return "IOC"
	case 0b1101:
		return "RTN"
	case 0b1110:
		return "SKZ"
	case 0b1111:
		return "NOPF"
	default:
		return ""
	}
}
