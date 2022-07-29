package main

import (
	"github.com/gin-gonic/gin"
	"github.com/notnil/chess"
	"hauthorn.me/chess-engine-in-go/engines"
	"net/http"
)

var game *chess.Game

var engine engines.ChessEngine

func main() {
	r := gin.Default()

	game = newGame()

	// engine = engines.RandomMoves{}
	engine = engines.NaiveEngine{Depth: 3}

	// All the static front end assets
	r.LoadHTMLFiles("web/index.html")
	r.Static("/css", "./web/css")
	r.Static("/img", "./web/img")
	r.Static("/js", "./web/js")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Chess engine requests
	r.POST("/reset", resetGame)
	r.POST("/move", playerMove)
	r.GET("/move", computerMove)

	// Start the engine
	r.Run(":8080")
}

type InputMove struct {
	San string
}

func resetGame(c *gin.Context) {
	game = newGame()
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
}

func playerMove(c *gin.Context) {
	var move InputMove
	// Bind/Map the move
	c.Bind(&move)

	// Check if the move is valid
	if err := game.MoveStr(move.San); err != nil {
		// handle error
		parsed := c.Error(err)
		c.JSON(http.StatusUnprocessableEntity, parsed.JSON())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
}

func computerMove(c *gin.Context) {
	move := engine.BestMove(game)

	if move == nil {
		c.JSON(http.StatusOK, gin.H{"move": nil, "fen": game.FEN()})
		return
	}
	game.Move(move)

	c.JSON(http.StatusOK, gin.H{"move": move.String(), "fen": game.FEN()})
}

func newGame() *chess.Game {
	return chess.NewGame()
}
