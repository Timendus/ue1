package ue1

import (
	"fmt"
	"strings"

	"github.com/timendus/ue1/helpers"
)

func Assemble(input string) ([]byte, error) {
	output := make([]byte, 0)

	lines := strings.Split(input, "\n")
	for num, line := range lines {
		// Ignore any comments
		commentStart := strings.Index(line, ";")
		if commentStart != -1 {
			line = line[:commentStart]
		}

		// If the line is empty, continue
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		line = strings.ToUpper(line)
		parts := strings.Fields(line)
		helpers.Assert(len(parts) > 0, "We should never get zero fields here")

		var bytecode byte

		switch parts[0] {
		case "NOP0":
			bytecode = 0
		case "LD":
			bytecode = 16
		case "ADD":
			bytecode = 32
		case "SUB":
			bytecode = 48
		case "ONE":
			bytecode = 64
		case "NAND":
			bytecode = 80
		case "OR":
			bytecode = 96
		case "XOR":
			bytecode = 112
		case "STO":
			bytecode = 128
		case "STOC":
			bytecode = 144
		case "IEN":
			bytecode = 160
		case "OEN":
			bytecode = 176
		case "IOC":
			bytecode = 192
		case "RTN":
			bytecode = 208
		case "SKZ":
			bytecode = 224
		case "NOPF":
			bytecode = 240
		default:
			return nil, fmt.Errorf("unknown opcode '%s' at line %d", parts[0], num+1)
		}

		if len(parts) > 1 {
			switch parts[1] {
			case "SR0":
				bytecode += 0
			case "SR1":
				bytecode += 1
			case "SR2":
				bytecode += 2
			case "SR3":
				bytecode += 3
			case "SR4":
				bytecode += 4
			case "SR5":
				bytecode += 5
			case "SR6":
				bytecode += 6
			case "SR7":
				bytecode += 7
			case "OR0", "RR":
				bytecode += 8
			case "OR1", "IR1":
				bytecode += 9
			case "OR2", "IR2":
				bytecode += 10
			case "OR3", "IR3":
				bytecode += 11
			case "OR4", "IR4":
				bytecode += 12
			case "OR5", "IR5":
				bytecode += 13
			case "OR6", "IR6":
				bytecode += 14
			case "OR7", "IR7":
				bytecode += 15
			default:
				return nil, fmt.Errorf("unknown operand '%s' at line %d", parts[1], num+1)
			}
		}

		output = append(output, bytecode)
	}

	return output, nil
}

func DisassembleInstruction(instruction byte) string {
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
