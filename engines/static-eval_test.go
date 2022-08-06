package engines

import (
	"github.com/notnil/chess"
	"testing"
)

func TestSimpleStaticScore(t *testing.T) {
	tests := []struct {
		name string
		fen  string
		want int
	}{
		{"Knight at corner is minus 20", "4k3/8/8/8/8/8/8/N3K3 w - - 0 1", 260},
		{"Knight at edge is minus 10", "4k3/8/8/7N/8/8/8/4K3 w - - 0 1", 270},
		{"Knight at center is plus 10", "4k3/8/8/4N3/8/8/8/4K3 w - - 0 1", 290},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FEN, _ := chess.FEN(tt.fen)
			game := chess.NewGame(FEN)
			position := game.Position()

			if got := SimpleStaticScore(position, position.Status()); got != tt.want {
				t.Errorf("SimpleStaticScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
