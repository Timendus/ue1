package main

type UE1 struct {
	pc    int   // Program counter
	sr    uint8 // Scratch register
	ir    uint8 // Input register (switches)
	or    uint8 // Output register
	rr    bool  // Result register
	ien   bool  // Input enable flag
	oen   bool  // Output enable flag
	carry bool  // Carry flag

	program []byte
	speed   int
	running bool
	state   int
}

const (
	STATE_RUNNING = iota
	STATE_HALTED
	STATE_STEP
)

func (cpu *UE1) step() {
	opcode := (cpu.program[cpu.pc] & 0xF0) >> 4
	operand := cpu.program[cpu.pc] & 0x0F
	cpu.pc = (cpu.pc + 1) % len(cpu.program)

	switch opcode {

	case 0b0000: // NOPO / NOP0

	case 0b0001: // LD
		cpu.rr = cpu.getValue(operand)

	case 0b0010: // ADD
		result := 0
		if cpu.rr {
			result += 1
		}
		if cpu.carry {
			result += 1
		}
		if cpu.getValue(operand) {
			result += 1
		}
		cpu.rr = (result & 1) != 0
		cpu.carry = (result & 2) != 0

	case 0b0011: // SUB
		result := 0
		if cpu.rr {
			result += 1
		}
		if cpu.carry {
			result += 1
		}
		if !cpu.getValue(operand) {
			result += 1
		}
		cpu.rr = (result & 1) != 0
		cpu.carry = (result & 2) != 0

	case 0b0100: // ONE
		cpu.rr = true
		cpu.carry = false

	case 0b0101: // NAND
		cpu.rr = !cpu.rr || !cpu.getValue(operand)

	case 0b0110: // OR
		cpu.rr = cpu.rr || cpu.getValue(operand)

	case 0b0111: // XOR
		cpu.rr = cpu.rr != cpu.getValue(operand)

	case 0b1000: // STO
		cpu.setValue(operand, cpu.rr)

	case 0b1001: // STOC
		cpu.setValue(operand, !cpu.rr)

	case 0b1010: // IEN
		cpu.ien = cpu.rr

	case 0b1011: // OEN
		cpu.oen = cpu.rr

	case 0b1100: // IOC
		bell()

	case 0b1101: // RTN
		cpu.pc += 1

	case 0b1110: // SKZ
		if !cpu.rr {
			cpu.pc += 1
		}

	case 0b1111: // NOPF
		cpu.state = STATE_HALTED

	default:
		fail("We should never see an unknown opcode")
	}
}

func (cpu *UE1) getValue(operand byte) bool {
	assert(operand < 16, "We should never see an operand over 15")

	if !cpu.ien {
		return false
	}
	if operand < 8 {
		return cpu.sr&(1<<operand) != 0
	}
	if operand == 8 {
		return cpu.rr
	}
	if operand < 16 {
		return cpu.ir&(1<<(operand-8)) != 0
	}
	return false
}

func (cpu *UE1) setValue(operand byte, value bool) {
	assert(operand < 16, "We should never see an operand over 15")

	if !cpu.oen {
		return
	}
	if operand < 8 {
		if value {
			cpu.sr = cpu.sr | (1 << operand)
		} else {
			cpu.sr = cpu.sr & ((1 << operand) ^ 0xff)
		}
		return
	}
	if operand < 16 {
		if value {
			cpu.or = cpu.or | (1 << (operand - 8))
		} else {
			cpu.or = cpu.or & ((1 << (operand - 8)) ^ 0xff)
		}
		return
	}
}
