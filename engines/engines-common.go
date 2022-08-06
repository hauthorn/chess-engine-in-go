package engines

import (
	"github.com/notnil/chess"
	"math"
)

type OrderedNumber interface {
	int64 | int | Alpha | Beta
}

type ChessEngine interface {
	BestMove(game *chess.Game) *chess.Move
}

var BigNumber = 100000

func IsGameOutcome(method chess.Method) bool {
	return method == chess.InsufficientMaterial ||
		method == chess.Stalemate ||
		method == chess.Checkmate
}

func SimpleStaticScore(position *chess.Position, method chess.Method) int {
	score := 0

	switch method {
	case chess.InsufficientMaterial:
		return 0
	case chess.Stalemate:
		return 0
	case chess.Checkmate:
		if position.Turn() == chess.White {
			return -BigNumber
		} else {
			return BigNumber
		}
	}

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
		if piece.Color() == chess.White {
			score += pieceValue
		} else {
			score -= pieceValue
		}
	}

	return score
}

func Abs[K OrderedNumber](x K) K {
	if x >= 0 {
		return x
	}

	return -x
}

func Max[K OrderedNumber](x K, y K) K {
	if x > y {
		return x
	}
	return y
}

func Min[K OrderedNumber](x K, y K) K {
	if x < y {
		return x
	}
	return y
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
