package engines

import (
	"github.com/notnil/chess"
	"reflect"
	"testing"
)

func TestAlphaBetaEngine_BestMove(t *testing.T) {
	t.Run("should try to capture", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 0,
		}
		g := chess.NewGame()
		g.MoveStr("d4")
		// moving knight out to capture d4 (we don't consider the counter by white
		g.MoveStr("Nc6")
		// Some pawn that cannot be captured
		g.MoveStr("c4")

		want := "c6d4"

		if got := r.BestMove(g); !reflect.DeepEqual(got.String(), want) {
			t.Errorf("BestMove() = %v, want %v", got, want)
		}
	})

	t.Run("should pick the capture that cannot be captured back", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 1,
		}
		g := chess.NewGame()
		g.MoveStr("d4")
		// moving knight out to capture either pawn
		g.MoveStr("Nc6")
		// move other pawn into place
		g.MoveStr("b4")

		want := "c6b4"
		if got := r.BestMove(g); !reflect.DeepEqual(got.String(), want) {
			t.Errorf("BestMove() = %v, want %v", got, want)
		}

		// Move the pawn up to defend, assume it will not try to capture any of them
		g.MoveStr("d6")
		g.MoveStr("c3")

		notWant := []string{"c6b4", "c6d4"}
		for _, not := range notWant {
			if got := r.BestMove(g); reflect.DeepEqual(got.String(), not) {
				t.Errorf("BestMove() = %v, shouldn't try to capture", got)
			}
		}
	})

	t.Run("should not try to capture a pawn that is defended", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 1,
		}
		g := chess.NewGame()
		g.MoveStr("d4")
		// moving knight out to capture either pawn
		g.MoveStr("Nc6")
		// move other pawn into place
		g.MoveStr("b4")
		// Move the pawn up to defend, assume it will not try to capture any of them
		g.MoveStr("d6")
		g.MoveStr("c3")

		notWant := []string{"c6b4", "c6d4"}
		for _, not := range notWant {
			if got := r.BestMove(g); reflect.DeepEqual(got.String(), not) {
				t.Errorf("BestMove() = %v, shouldn't try to capture", got)
			}
		}
	})

	t.Run("should go for mate", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 0,
		}

		// About to mate the black king with king and tower, but can also choose to capture rook
		fen, _ := chess.FEN("3k4/6Rr/3K4/8/8/8/8/8 w - - 0 1")
		g := chess.NewGame(fen)

		// Move the rook up to mate
		want := "g7g8"
		if got := r.BestMove(g); !reflect.DeepEqual(got.String(), want) {
			t.Errorf("BestMove() = %v, want %v", got, want)
		}
	})

	t.Run("should go for mate in 2", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 1,
		}

		// About to mate the black king with king and tower, but can also choose to capture rook
		fen, _ := chess.FEN("k7/6Rr/2K5/7n/7p/8/8/8 w - - 0 1")
		g := chess.NewGame(fen)

		// Capture rook
		want := "g7h7"
		if got := r.BestMove(g); !reflect.DeepEqual(got.String(), want) {
			t.Errorf("BestMove() = %v, want %v", got, want)
		}
	})

	t.Run("should prefer pawn advancement in center", func(t *testing.T) {
		r := AlphaBetaEngine{
			Depth: 0,
		}

		// Just a single pawn, should prefer to advance it 2 spaces in the center
		fen, _ := chess.FEN("4k3/8/8/8/8/8/P3P2P/4K3 w - - 0 1")
		g := chess.NewGame(fen)

		want := "e2e4"
		if got := r.BestMove(g); !reflect.DeepEqual(got.String(), want) {
			t.Errorf("BestMove() = %v, want %v", got, want)
		}
	})
}
