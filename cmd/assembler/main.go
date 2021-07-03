package main

import (
	"fmt"
	"os"

	"github.com/hiragi-gkuth/n2t_assembler/pkg/assembler"
)

func main() {
	filePath := os.Args[1]
	assembler := assembler.NewAssembler()

	parser := assembler.Load(filePath)
	codes := assembler.Assemble(parser)

	for i, code := range codes {
		fmt.Printf("%d: %016b\n", i, code)
	}
}
