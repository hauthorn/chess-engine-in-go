package engines

import (
	"github.com/notnil/chess"
)

type NaiveEngine struct {
	Depth uint8
}

func (r NaiveEngine) BestMove(game *chess.Game) *chess.Move {
	validMoves := game.ValidMoves()

	bestScore := -100000
	var bestMove *chess.Move = nil

	// For every valid move, calculate the score of the move
	for _, m := range validMoves {
		g := game.Clone()
		g.Move(m)

		score := score(g, r.Depth)

		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	return bestMove
}

func score(game *chess.Game, depth uint8) int {
	position := game.Position()
	currentTurn := position.Turn()

	switch game.Outcome() {
	case chess.Draw:
		return 0
	case chess.WhiteWon:
		if currentTurn == chess.Black {
			return 1000000
		}
		return -1000000
	case chess.BlackWon:
		if currentTurn == chess.White {
			return 1000000
		}
		return -1000000
	}

	if depth == 0 {
		score := 0

		for i := chess.Square(0); i < 64; i++ {
			piece := position.Board().Piece(i)
			pieceValue := 0
			switch piece.Type() {
			case chess.King:
				pieceValue = 100000
			case chess.Queen:
				pieceValue = 900
			case chess.Rook:
				pieceValue = 500
			case chess.Bishop:
				pieceValue = 320
			case chess.Knight:
				pieceValue = 300
			case chess.Pawn:
				pieceValue = 100
			}
			if piece.Color() != currentTurn {
				score += pieceValue
			} else {
				score -= pieceValue
			}
		}

		return score
	}

	worstScore := 100000

	// For every valid move, calculate the score of the move
	for _, m := range game.ValidMoves() {
		g := game.Clone()
		g.Move(m)

		score := -score(g, depth-1)

		if score < worstScore {
			worstScore = score
		}
	}
	return worstScore
}
