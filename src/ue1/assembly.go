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

		// Split the line into uppercase words, of which we will only use two
		line = strings.ToUpper(line)
		parts := strings.Fields(line)
		helpers.Assert(len(parts) > 0, "We should never get zero fields here")

		// Find the opcode the user wants
		opcode, ok := opcodeMap[parts[0]]
		if !ok {
			return nil, fmt.Errorf("unrecognized opcode '%s' on line %d", parts[0], num)
		}

		// Find the optional operand the user wants
		var operand Operand
		if len(parts) > 1 {
			operand, ok = operandMap[parts[1]]
			if !ok {
				return nil, fmt.Errorf("unrecognized operand '%s' to opcode '%s' on line %d", parts[1], parts[0], num)
			}
		}

		// Combine the opcode and operand and add to the binary
		output = append(output, byte(opcode)<<4+byte(operand))
	}

	return output, nil
}

func DisassembleInstruction(instruction byte) string {
	opcode := Opcode(instruction & 0xF0 >> 4)
	operand := Operand(instruction & 0xF)

	switch opcode {
	case LD, ADD, SUB, NAND, OR, XOR:
		return opcodeList[opcode] + " " + inputOperands[operand]
	case STO, STOC:
		return opcodeList[opcode] + " " + outputOperands[operand]
	default:
		return opcodeList[opcode]
	}
}
