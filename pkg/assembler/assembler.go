package assembler

import (
	"strconv"
)

type BinaryCode []uint16

type IAssembler interface {
	Load(filePath string) IParser
	Assemble(parser IParser) BinaryCode
}

type Assembler struct{}

func NewAssembler() IAssembler {
	return Assembler{}
}

func (a Assembler) Load(filePath string) IParser {
	return NewParser(filePath)
}

func (a Assembler) Assemble(parser IParser) BinaryCode {
	// PASS 1: resolve asm code labels
	symbolTable := symbolLabelResolver(parser)
	// PASS 2: let asm commands to hack code
	hackBinaryCode := assemble(parser, symbolTable)

	return hackBinaryCode
}

func symbolLabelResolver(parser IParser) ISymbolTable {
	symbolTable := NewSymbolTable()
	romAddr := uint16(0x0)
	for {
		cType := parser.CommandType()

		if cType == A_COMMAND || cType == C_COMMAND {
			romAddr++
		}
		if cType == L_COMMAND {
			symbolTable.AddLabelEntry(parser.Symbol(), romAddr)
		}

		if !parser.HasMoreCommands() {
			break
		}
		parser.Advance()
	}
	return symbolTable
}

func assemble(parser IParser, symbolTable ISymbolTable) []uint16 {
	parser.ResetSeeker()
	coder := NewCode()
	hackBinaryCode := []uint16{}

	for {
		cType := parser.CommandType()
		// do not assemble L_COMMAND on pass 2
		if cType == L_COMMAND {
			if !parser.HasMoreCommands() {
				break
			}
			parser.Advance()
			continue
		}

		var commandBits uint16
		switch {
		case cType == C_COMMAND:
			commandBits = assembleCCommand(
				coder.Dest(parser.Dest()),
				coder.Jump(parser.Jump()),
				coder.Comp(parser.Comp()),
			)
		case cType == A_COMMAND:
			commandBits = assembleACommand(symbolTable, parser.Symbol())
		default:
			panic("unrecognized commands")
		}

		hackBinaryCode = append(hackBinaryCode, commandBits)
		if !parser.HasMoreCommands() {
			break
		}
		parser.Advance()
	}
	return hackBinaryCode
}

func assembleCCommand(dest, jump, comp uint16) (commandBits uint16) {
	commandBitsBase := uint16(0b1110_0000_0000_0000)
	commandBits = commandBitsBase | jump | (dest << 3) | (comp << 6)
	return
}

func assembleACommand(symbolTable ISymbolTable, symbol string) (commandBits uint16) {
	symbolValue, e := strconv.Atoi(symbol)
	if e == nil {
		// use symbol value as commandBits if symbol is immediate value
		commandBits = uint16(symbolValue)
		return
	}

	// symbol resolver
	if !symbolTable.Contains(symbol) {
		symbolTable.AddVariableEntry(symbol)
	}
	commandBits = symbolTable.GetAddress(symbol)
	return
}
