package assembler_test

import (
	"testing"

	"github.com/hiragi-gkuth/n2t_assembler/pkg/assembler"
)

func BenchmarkCodeDest(b *testing.B) {
	code := assembler.NewCode()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code.Dest("AMD")
	}
}

func BenchmarkCodeJump(b *testing.B) {
	code := assembler.NewCode()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code.Jump("JLE")
	}
}

func BenchmarkCodeComp(b *testing.B) {
	code := assembler.NewCode()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code.Comp("M-D")
	}
}
