package engines

import (
	"github.com/notnil/chess"
)

type AlphaBetaEngine struct {
	Depth uint8
}

type Alpha int
type Beta int

func (e AlphaBetaEngine) BestMove(game *chess.Game) *chess.Move {
	alpha := Alpha(-BigNumber)
	beta := Beta(BigNumber)
	var bestMove *chess.Move = nil
	var bestScore int
	position := game.Position()
	turn := position.Turn()
	if turn == chess.White {
		bestScore = -BigNumber
	} else {
		bestScore = BigNumber
	}

	for _, move := range position.ValidMoves() {
		newPos := position.Update(move)
		score := alphaBetaScore(newPos, e.Depth, alpha, beta)

		if turn == chess.White {
			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		} else {
			if score < bestScore {
				bestScore = score
				bestMove = move
			}
		}

	}

	return bestMove
}

func alphaBetaScore(position *chess.Position, depth uint8, alpha Alpha, beta Beta) int {
	method := position.Status()
	if depth == 0 || IsGameOutcome(method) {
		return SimpleStaticScore(position, method)
	}

	// Maximize
	if position.Turn() == chess.White {
		maxScore := -BigNumber
		validMoves := position.ValidMoves()

		for _, move := range validMoves {
			score := alphaBetaScore(position.Update(move), depth-1, alpha, beta)
			maxScore = Max(score, maxScore)
			alpha = Max(Alpha(score), alpha)
			if int(beta) <= int(alpha) {
				break
			}
		}

		return maxScore
	} else {
		// Minimize
		minScore := BigNumber
		validMoves := position.ValidMoves()
		for _, move := range validMoves {
			score := alphaBetaScore(position.Update(move), depth-1, alpha, beta)
			minScore = Min(score, minScore)
			beta = Min(Beta(score), beta)
			if int(beta) <= int(alpha) {
				break
			}
		}

		return minScore
	}
}
