package engines

import "github.com/notnil/chess"

type ChessEngine interface {
	BestMove(game *chess.Game) *chess.Move
}

func SimpleStaticScore(position *chess.Position) int {
	score := 0

	for i := chess.Square(0); i < 64; i++ {
		piece := position.Board().Piece(i)
		pieceValue := 0
		switch piece.Type() {
		case chess.Queen:
			pieceValue = 900
		case chess.Rook:
			pieceValue = 500
		case chess.Bishop:
			pieceValue = 320
		case chess.Knight:
			pieceValue = 300
		case chess.Pawn:
			pieceValue = pawnPoints(i, piece.Color())
		}
		if piece.Color() != position.Turn() {
			score += pieceValue
		} else {
			score -= pieceValue
		}
	}

	return score
}
