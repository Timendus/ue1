package ue1

type Opcode byte
type Operand byte

const (
	NOP0 Opcode = 0b0000
	LD   Opcode = 0b0001
	ADD  Opcode = 0b0010
	SUB  Opcode = 0b0011
	ONE  Opcode = 0b0100
	NAND Opcode = 0b0101
	OR   Opcode = 0b0110
	XOR  Opcode = 0b0111
	STO  Opcode = 0b1000
	STOC Opcode = 0b1001
	IEN  Opcode = 0b1010
	OEN  Opcode = 0b1011
	IOC  Opcode = 0b1100
	RTN  Opcode = 0b1101
	SKZ  Opcode = 0b1110
	NOPF Opcode = 0b1111

	SR0 Operand = 0b0000
	SR1 Operand = 0b0001
	SR2 Operand = 0b0010
	SR3 Operand = 0b0011
	SR4 Operand = 0b0100
	SR5 Operand = 0b0101
	SR6 Operand = 0b0110
	SR7 Operand = 0b0111
	OR0 Operand = 0b1000
	OR1 Operand = 0b1001
	OR2 Operand = 0b1010
	OR3 Operand = 0b1011
	OR4 Operand = 0b1100
	OR5 Operand = 0b1101
	OR6 Operand = 0b1110
	OR7 Operand = 0b1111
)

var opcodeMap = map[string]Opcode{
	"NOP0": NOP0,
	"LD":   LD,
	"ADD":  ADD,
	"SUB":  SUB,
	"ONE":  ONE,
	"NAND": NAND,
	"OR":   OR,
	"XOR":  XOR,
	"STO":  STO,
	"STOC": STOC,
	"IEN":  IEN,
	"OEN":  OEN,
	"IOC":  IOC,
	"RTN":  RTN,
	"SKZ":  SKZ,
	"NOPF": NOPF,
}

var operandMap = map[string]Operand{
	"SR0": SR0,
	"SR1": SR1,
	"SR2": SR2,
	"SR3": SR3,
	"SR4": SR4,
	"SR5": SR5,
	"SR6": SR6,
	"SR7": SR7,

	"RR":  OR0,
	"IR1": OR1,
	"IR2": OR2,
	"IR3": OR3,
	"IR4": OR4,
	"IR5": OR5,
	"IR6": OR6,
	"IR7": OR7,

	"OR0": OR0,
	"OR1": OR1,
	"OR2": OR2,
	"OR3": OR3,
	"OR4": OR4,
	"OR5": OR5,
	"OR6": OR6,
	"OR7": OR7,
}

// Mapping from integer to string for opcodes
var opcodeList = [16]string{
	"NOP0",
	"LD",
	"ADD",
	"SUB",
	"ONE",
	"NAND",
	"OR",
	"XOR",
	"STO",
	"STOC",
	"IEN",
	"OEN",
	"IOC",
	"RTN",
	"SKZ",
	"NOPF",
}

// Mapping from integer to string for input operands
var inputOperands = [16]string{
	"SR0",
	"SR1",
	"SR2",
	"SR3",
	"SR4",
	"SR5",
	"SR6",
	"SR7",
	"RR",
	"IR1",
	"IR2",
	"IR3",
	"IR4",
	"IR5",
	"IR6",
	"IR7",
}

// Mapping from integer to string for output operands
var outputOperands = [16]string{
	"SR0",
	"SR1",
	"SR2",
	"SR3",
	"SR4",
	"SR5",
	"SR6",
	"SR7",
	"OR0",
	"OR1",
	"OR2",
	"OR3",
	"OR4",
	"OR5",
	"OR6",
	"OR7",
}
