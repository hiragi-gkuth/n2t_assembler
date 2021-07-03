package assembler_test

import (
	"testing"

	"github.com/hiragi-gkuth/n2t_assembler/pkg/assembler"
)

func BenchmarkAssembler(b *testing.B) {
	as := assembler.NewAssembler()
	pa := as.Load("/Users/hiragi-gkuth/go/src/github.com/hiragi-gkuth/n2t_assembler/test/PongL.asm")
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		as.Assemble(pa)
	}
	b.StopTimer()
}
