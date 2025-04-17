package ue1

import (
	"github.com/timendus/ue1/helpers"
)

type UE1 struct {
	PC    int   // Program counter
	SR    uint8 // Scratch register
	IR    uint8 // Input register (switches)
	OR    uint8 // Output register
	RR    bool  // Result register
	IEN   bool  // Input enable flag
	OEN   bool  // Output enable flag
	Carry bool  // Carry flag

	Program  []byte
	Speed    int
	Running  bool
	State    int
	BellFunc func()
}

const (
	STATE_RUNNING = iota
	STATE_HALTED
	STATE_STEP
)

func (cpu *UE1) Step() {
	opcode := (cpu.Program[cpu.PC] & 0xF0) >> 4
	operand := cpu.Program[cpu.PC] & 0x0F
	cpu.PC = (cpu.PC + 1) % len(cpu.Program)

	switch opcode {

	case 0b0000: // NOPO / NOP0

	case 0b0001: // LD
		cpu.RR = cpu.getValue(operand)

	case 0b0010: // ADD
		result := 0
		if cpu.RR {
			result += 1
		}
		if cpu.Carry {
			result += 1
		}
		if cpu.getValue(operand) {
			result += 1
		}
		cpu.RR = (result & 1) != 0
		cpu.Carry = (result & 2) != 0

	case 0b0011: // SUB
		result := 0
		if cpu.RR {
			result += 1
		}
		if cpu.Carry {
			result += 1
		}
		if !cpu.getValue(operand) {
			result += 1
		}
		cpu.RR = (result & 1) != 0
		cpu.Carry = (result & 2) != 0

	case 0b0100: // ONE
		cpu.RR = true
		cpu.Carry = false

	case 0b0101: // NAND
		cpu.RR = !cpu.RR || !cpu.getValue(operand)

	case 0b0110: // OR
		cpu.RR = cpu.RR || cpu.getValue(operand)

	case 0b0111: // XOR
		cpu.RR = cpu.RR != cpu.getValue(operand)

	case 0b1000: // STO
		cpu.setValue(operand, cpu.RR)

	case 0b1001: // STOC
		cpu.setValue(operand, !cpu.RR)

	case 0b1010: // IEN
		cpu.IEN = cpu.RR

	case 0b1011: // OEN
		cpu.OEN = cpu.RR

	case 0b1100: // IOC
		cpu.BellFunc()

	case 0b1101: // RTN
		cpu.PC += 1

	case 0b1110: // SKZ
		if !cpu.RR {
			cpu.PC += 1
		}

	case 0b1111: // NOPF
		cpu.State = STATE_HALTED

	default:
		helpers.Fail("We should never see an unknown opcode")
	}
}

func (cpu *UE1) getValue(operand byte) bool {
	helpers.Assert(operand < 16, "We should never see an operand over 15")

	if !cpu.IEN {
		return false
	}
	if operand < 8 {
		return cpu.SR&(1<<operand) != 0
	}
	if operand == 8 {
		return cpu.RR
	}
	if operand < 16 {
		return cpu.IR&(1<<(operand-8)) != 0
	}
	return false
}

func (cpu *UE1) setValue(operand byte, value bool) {
	helpers.Assert(operand < 16, "We should never see an operand over 15")

	if !cpu.OEN {
		return
	}
	if operand < 8 {
		if value {
			cpu.SR = cpu.SR | (1 << operand)
		} else {
			cpu.SR = cpu.SR & ((1 << operand) ^ 0xff)
		}
		return
	}
	if operand < 16 {
		if value {
			cpu.OR = cpu.OR | (1 << (operand - 8))
		} else {
			cpu.OR = cpu.OR & ((1 << (operand - 8)) ^ 0xff)
		}
		return
	}
}
