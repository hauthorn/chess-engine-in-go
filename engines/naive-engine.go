package engines

import (
	"github.com/notnil/chess"
	"math"
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
		sourcePosition := game.Position()
		afterMove := sourcePosition.Update(m)
		score := score(afterMove, r.Depth)

		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	return bestMove
}

func score(position *chess.Position, depth uint8) int {
	currentColor := position.Turn()
	outcome := position.Status()

	switch outcome {
	case chess.InsufficientMaterial:
		return 0
	case chess.Stalemate:
		return 0
	case chess.Checkmate:
		return 1000000
	}

	if depth == 0 {
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
			if piece.Color() != currentColor {
				score += pieceValue
			} else {
				score -= pieceValue
			}
		}

		return score
	}

	worstScore := 100000

	// For every valid move, calculate the score of the move
	for _, m := range position.ValidMoves() {
		afterMove := position.Update(m)

		score := -score(afterMove, depth-1)

		if score < worstScore {
			worstScore = score
		}
	}
	return worstScore
}

func pawnPoints(i chess.Square, color chess.Color) int {
	rankPoints := 0
	if color == chess.White {
		rankPoints = int(i.Rank()) - 1
	} else {
		rankPoints = 7 - int(i.Rank())
	}

	file := i.File()
	distance := math.Abs(3.5 - float64(file))
	filePoints := 3 - int(math.Floor(distance))

	score := 100 + rankPoints*5 + filePoints*rankPoints
	return score
}
