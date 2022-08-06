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

// AlphaBetaEngine benchmarks

func BenchmarkAlphaBetaEngine_BestMoveDepth2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := AlphaBetaEngine{Depth: 2}
		naive.BestMove(g)
	}
}

func BenchmarkAlphaBetaEngine_BestMoveDepth3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := AlphaBetaEngine{Depth: 3}
		naive.BestMove(g)
	}
}

func BenchmarkAlphaBetaEngine_BestMoveDepth4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := AlphaBetaEngine{Depth: 4}
		naive.BestMove(g)
	}
}

func BenchmarkAlphaBetaEngine_BestMoveDepth5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := chess.NewGame()
		naive := AlphaBetaEngine{Depth: 5}
		naive.BestMove(g)
	}
}
