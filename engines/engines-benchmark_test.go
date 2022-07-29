package engines

import (
	"github.com/notnil/chess"
	"testing"
)

// NaiveEngine benchmarks

func BenchmarkNaiveEngine_BestMoveDepth2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := NaiveEngine{Depth: 2}
		naive.BestMove(g)
	}
}

func BenchmarkNaiveEngine_BestMoveDepth3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := NaiveEngine{Depth: 3}
		naive.BestMove(g)
	}
}

func BenchmarkNaiveEngine_BestMoveDepth4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := NaiveEngine{Depth: 4}
		naive.BestMove(g)
	}
}
