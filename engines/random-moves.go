package engines

import (
	"github.com/notnil/chess"
	"math/rand"
)

type RandomMoves struct {
}

func (r RandomMoves) BestMove(game *chess.Game) *chess.Move {
	validMoves := game.ValidMoves()

	if len(validMoves) == 0 {
		return nil
	}

	index := rand.Intn(len(validMoves))

	return validMoves[index]
}
