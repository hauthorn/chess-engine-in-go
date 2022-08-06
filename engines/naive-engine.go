package engines

import (
	"github.com/notnil/chess"
)

type NaiveEngine struct {
	Depth uint8
}

func (r NaiveEngine) BestMove(game *chess.Game) *chess.Move {
	validMoves := game.ValidMoves()

	var bestScore int
	turn := game.Position().Turn()
	if turn == chess.White {
		bestScore = -BigNumber
	} else {
		bestScore = BigNumber
	}
	var bestMove *chess.Move = nil

	// For every valid move, calculate the score of the move
	for _, m := range validMoves {
		sourcePosition := game.Position()
		afterMove := sourcePosition.Update(m)
		score := score(afterMove, r.Depth)

		if turn == chess.White {
			if score > bestScore {
				bestScore = score
				bestMove = m
			}
		} else {
			if score < bestScore {
				bestScore = score
				bestMove = m
			}
		}
	}

	return bestMove
}

func score(position *chess.Position, depth uint8) int {
	turn := position.Turn()

	method := position.Status()
	if depth == 0 || IsGameOutcome(method) {
		return SimpleStaticScore(position, method)
	}

	var worstScore = 0
	if turn == chess.White {
		worstScore = -BigNumber
	} else {
		worstScore = BigNumber
	}

	// For every valid move, calculate the score of the move
	for _, m := range position.ValidMoves() {
		afterMove := position.Update(m)

		score := score(afterMove, depth-1)

		if turn == chess.White {
			worstScore = Max(worstScore, score)
		} else {
			worstScore = Min(worstScore, score)
		}
	}
	return worstScore
}
