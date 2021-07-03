package assembler

import "strings"

type ICode interface {
	Dest(mnemonic string) uint16
	Comp(nmemonic string) uint16
	Jump(nmemonic string) uint16
}

type Code struct{}

var (
	compInst = map[string]uint16{
		"0":   0b0_101_010,
		"1":   0b0_111_111,
		"-1":  0b0_111_010,
		"D":   0b0_001_100,
		"A":   0b0_110_000,
		"!D":  0b0_001_101,
		"!A":  0b0_110_000,
		"-D":  0b0_001_111,
		"-A":  0b0_110_011,
		"D+1": 0b0_011_111,
		"A+1": 0b0_110_111,
		"D-1": 0b0_001_110,
		"A-1": 0b0_110_010,
		"D+A": 0b0_000_010,
		"D-A": 0b0_010_011,
		"A-D": 0b0_000_111,
		"D&A": 0b0_000_000,
		"D|A": 0b0_010_101,
		"M":   0b1_110_000,
		"!M":  0b1_110_001,
		"-M":  0b1_110_011,
		"M+1": 0b1_110_111,
		"M-1": 0b1_110_010,
		"D+M": 0b1_000_010,
		"D-M": 0b1_010_011,
		"M-D": 0b1_000_111,
		"D&M": 0b1_000_000,
		"D|M": 0b1_010_101,
	}
)

func NewCode() ICode {
	return &Code{}
}

func (c Code) Dest(mnemonic string) uint16 {
	bits := 0b000

	if strings.Contains(mnemonic, "A") {
		bits |= 0b100
	}
	if strings.Contains(mnemonic, "D") {
		bits |= 0b010
	}
	if strings.Contains(mnemonic, "M") {
		bits |= 0b001
	}
	return uint16(bits)
}

func (c Code) Jump(mnemonic string) uint16 {
	switch mnemonic {
	case "":
		return 0b000
	case "JGT":
		return 0b001
	case "JEQ":
		return 0b010
	case "JGE":
		return 0b011
	case "JLT":
		return 0b100
	case "JNE":
		return 0b101
	case "JLE":
		return 0b110
	case "JMP":
		return 0b111
	default:
		panic("Invalid mnemonic is taken")
	}
}

func (c Code) Comp(mnemonic string) uint16 {
	bits, ok := compInst[mnemonic]
	if !ok {
		panic("Invalid mnemonic is taken")
	}
	return bits
}
