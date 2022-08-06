package engines

import (
	"github.com/notnil/chess"
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
