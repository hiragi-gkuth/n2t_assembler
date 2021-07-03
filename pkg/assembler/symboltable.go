package assembler

import (
	"fmt"
	"strconv"
)

type ISymbolTable interface {
	AddLabelEntry(symbol string, romAddress uint16)
	AddVariableEntry(symbol string)
	Contains(symbol string) bool
	GetAddress(symbol string) uint16
	Show()
}

type SymbolTable struct {
	table             map[string]uint16
	variableAddrIndex uint16
}

func NewSymbolTable() ISymbolTable {

	// Hack machine defined symbols
	table := map[string]uint16{
		"SP":     0x0000,
		"LCL":    0x0001,
		"ARG":    0x0002,
		"THIS":   0x0003,
		"THAT":   0x0004,
		"SCREEN": 0x4000,
		"KBD":    0x6000,
	}
	for rx := 0; rx < 16; rx++ {
		table["R"+strconv.Itoa(rx)] = uint16(rx)
	}
	return &SymbolTable{
		table:             table,
		variableAddrIndex: 0x0010,
	}
}

func (s SymbolTable) Show() {
	for k, v := range s.table {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func (s *SymbolTable) AddLabelEntry(symbol string, address uint16) {
	s.table[symbol] = address
}

func (s *SymbolTable) AddVariableEntry(symbol string) {
	s.table[symbol] = s.variableAddrIndex
	s.variableAddrIndex++
}

func (s SymbolTable) Contains(symbol string) bool {
	_, contain := s.table[symbol]
	return contain
}

func (s SymbolTable) GetAddress(symbol string) uint16 {
	address, ok := s.table[symbol]
	if !ok {
		panic(fmt.Sprintf("undefined symbol for %s", symbol))
	}
	return address
}
