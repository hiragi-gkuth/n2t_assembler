package assembler_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/hiragi-gkuth/n2t_assembler/pkg/assembler"
)

func BenchmarkParserAdvance(b *testing.B) {
	parser := assembler.NewParser("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		parser.Advance()
	}
}

func BenchmarkParserResetSeeker(b *testing.B) {
	parser := assembler.NewParser("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		parser.ResetSeeker()
	}
}

func BenchmarkParserCommandTypeAndSymbol(b *testing.B) {
	parser := assembler.NewParser("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !parser.HasMoreCommands() {
			parser.ResetSeeker()
		}
		cType := parser.CommandType()

		if cType == assembler.C_COMMAND {
			parser.Jump()
			parser.Comp()
			parser.Dest()
		}
		if cType == assembler.L_COMMAND {
			parser.Symbol()
		}
	}
}

func BenchmarkParserLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, e := ioutil.ReadFile("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")
		if e != nil {
			panic(e)
		}
	}
}

func BenchmarkParserAsmRawCodeToCommands(b *testing.B) {
	rawAsmCode, e := ioutil.ReadFile("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")
	if e != nil {
		panic(e)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		asmRawCodeToCommands(rawAsmCode)
	}
}

func asmRawCodeToCommands(rawCode []byte) (commands []string) {
	rawCodeArray := strings.Split(string(rawCode), "\n")
	commands = make([]string, 0, len(rawCodeArray))
	// skipping comments and empty line
	for _, line := range rawCodeArray {
		line = strings.TrimSpace(line)

		// skip comment
		commentPos := strings.Index(line, "//")
		if commentPos != -1 {
			line = line[:commentPos]
		}

		// skip empty line
		if len(line) == 0 {
			continue
		}
		commands = append(commands, line)
	}
	return
}
