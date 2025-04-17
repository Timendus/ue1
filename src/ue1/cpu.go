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
	State    int
	BellFunc func()
}

const (
	STATE_RUNNING = iota
	STATE_HALTED
	STATE_STEP
)

func (cpu *UE1) Step() {
	opcode := Opcode((cpu.Program[cpu.PC] & 0xF0) >> 4)
	operand := Operand(cpu.Program[cpu.PC] & 0x0F)
	cpu.PC = (cpu.PC + 1) % len(cpu.Program)

	switch opcode {

	case NOP0:
		// No operation

	case LD:
		cpu.RR = cpu.getValue(operand)

	case ADD:
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

	case SUB:
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

	case ONE:
		cpu.RR = true
		cpu.Carry = false

	case NAND:
		cpu.RR = !cpu.RR || !cpu.getValue(operand)

	case OR:
		cpu.RR = cpu.RR || cpu.getValue(operand)

	case XOR:
		cpu.RR = cpu.RR != cpu.getValue(operand)

	case STO:
		cpu.setValue(operand, cpu.RR)

	case STOC:
		cpu.setValue(operand, !cpu.RR)

	case IEN:
		cpu.IEN = cpu.RR

	case OEN:
		cpu.OEN = cpu.RR

	case IOC:
		cpu.BellFunc()

	case RTN:
		cpu.PC += 1

	case SKZ:
		if !cpu.RR {
			cpu.PC += 1
		}

	case NOPF:
		cpu.State = STATE_HALTED

	default:
		helpers.Fail("We should never see an unknown opcode")
	}
}

func (cpu *UE1) getValue(operand Operand) bool {
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

func (cpu *UE1) setValue(operand Operand, value bool) {
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
