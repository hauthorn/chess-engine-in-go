package engines

import (
	"github.com/notnil/chess"
	"math"
)

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
			pieceValue = bishopPoints(i, piece.Color())
		case chess.Knight:
			pieceValue = knightPoints(i)
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

	score := 100 + rankPoints*3 + filePoints*5
	return score
}

func knightPoints(i chess.Square) int {
	baseScore := 280

	// Knights should avoid the edges of the board
	positionScore := 0
	rank := i.Rank()
	if rank == 0 || rank == 7 {
		positionScore -= 10
	}
	file := i.File()
	if file == 0 || file == 7 {
		positionScore -= 10
	}

	// Reward centered knights
	if rank < 5 && rank > 2 && file < 5 && file > 2 {
		positionScore += 10
	}

	return baseScore + positionScore
}

func bishopPoints(i chess.Square, color chess.Color) int {
	baseScore := 320
	positionScore := 0

	// We would like to punish bishops still in their starting location
	if color == chess.White && (i == chess.F1 || i == chess.C1) {
		positionScore -= 10
	}
	if color == chess.Black && (i == chess.F8 || i == chess.C8) {
		positionScore -= 10
	}

	return baseScore + positionScore
}
