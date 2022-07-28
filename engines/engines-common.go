package engines

import "github.com/notnil/chess"

type ChessEngine interface {
	BestMove(game *chess.Game) *chess.Move
}
