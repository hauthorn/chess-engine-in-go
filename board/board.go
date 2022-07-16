package board

import (
	"strconv"
	"strings"
)

type Color bool
type PieceType uint8

const BLACK Color = false
const WHITE Color = true

const (
	Empty  PieceType = iota
	Pawn   PieceType = iota
	Knight PieceType = iota
	Bishop PieceType = iota
	Rook   PieceType = iota
	Queen  PieceType = iota
	King   PieceType = iota
)

type Piece struct {
	Color Color
	Type  PieceType
}

func (p Piece) String() string {
	if p.Type == Empty {
		return "E"
	}

	var pieceType = "EM"
	switch p.Type {
	case Pawn:
		pieceType = "P"
	case Knight:
		pieceType = "N"
	case Bishop:
		pieceType = "B"
	case Rook:
		pieceType = "R"
	case Queen:
		pieceType = "Q"
	case King:
		pieceType = "K"
	}

	if !p.Color {
		return strings.ToLower(pieceType)
	}

	return pieceType
}

// EM No piece present here
var EM = Piece{BLACK, Empty}

// BR Rooks
var BR = Piece{Color: BLACK, Type: Rook}
var WR = Piece{Color: WHITE, Type: Rook}

// BN Knights
var BN = Piece{Color: BLACK, Type: Knight}
var WN = Piece{Color: WHITE, Type: Knight}

// BB Bishops
var BB = Piece{Color: BLACK, Type: Bishop}
var WB = Piece{Color: WHITE, Type: Bishop}

// BQ Queens
var BQ = Piece{Color: BLACK, Type: Queen}
var WQ = Piece{Color: WHITE, Type: Queen}

// BK Kings
var BK = Piece{Color: BLACK, Type: King}
var WK = Piece{Color: WHITE, Type: King}

// BP Pawns
var BP = Piece{Color: BLACK, Type: Pawn}
var WP = Piece{Color: WHITE, Type: Pawn}

type File uint8
type Rank uint8

const (
	A File = 0
	B File = 1
	C File = 2
	D File = 3
	E File = 4
	F File = 5
	G File = 6
	H File = 7
)

type Position struct {
	File File
	Rank Rank
}

func (p Position) String() string {
	fileString := ""
	switch p.File {
	case A:
		fileString = "A"
	case B:
		fileString = "B"
	case C:
		fileString = "C"
	case D:
		fileString = "D"
	case E:
		fileString = "E"
	case F:
		fileString = "F"
	case G:
		fileString = "G"
	case H:
		fileString = "H"
	}
	return fileString + strconv.Itoa(int(p.Rank))
}

type Board struct {
	Positions [8][8]Piece
	ToMove    Color
}

// DefaultBoard is the default starting position in a game of Chess
func DefaultBoard() Board {
	board := Board{
		Positions: [8][8]Piece{
			{WR, WN, WB, WQ, WK, WB, WN, WR},
			{WP, WP, WP, WP, WP, WP, WP, WP},
			{},
			{},
			{},
			{},
			{BP, BP, BP, BP, BP, BP, BP, BP},
			{BR, BN, BB, BQ, BK, BB, BN, BR},
		},
		ToMove: WHITE,
	}
	return board
}

// KingsAndRook Board with just 2 kings and a Rook, limiting some black King moves
func KingsAndRook(toMove Color) Board {
	board := Board{}
	board.Positions = [8][8]Piece{
		{},
		{EM, EM, EM, WK, WR, EM, EM, EM},
		{},
		{},
		{},
		{},
		{},
		{EM, EM, EM, BK, EM, EM, EM, EM},
	}
	board.ToMove = toMove
	return board
}

func (b Board) getPieceAt(position Position) *Piece {
	return &b.Positions[position.Rank-1][position.File]
}

func (b Board) Clone() Board {
	positions := [8][8]Piece{}

	for file := File(0); file < 8; file++ {
		for rank := Rank(0); rank < 8; rank++ {
			positions[rank][file] = b.p(rank, file)
		}
	}
	board := Board{
		positions,
		b.ToMove,
	}
	return board
}

func (b Board) FENString() string {
	// Piece Placement
	piecePlacement := ""
	ranks := make([]string, 0)
	for rank := Rank(7); rank+1 > 0; rank-- {
		emptyCount := 0
		rankString := ""

		for file := File(0); file <= 7; file++ {
			piece := b.p(rank, file)
			switch piece.Type {
			case Empty:
				emptyCount++
			default:
				if emptyCount > 0 {
					rankString = rankString + strconv.Itoa(emptyCount)
					emptyCount = 0
				}
				rankString = rankString + piece.String()
			}
			if file == 7 && emptyCount > 0 {
				rankString = rankString + strconv.Itoa(emptyCount)
				emptyCount = 0
			}
		}

		ranks = append(ranks, rankString)
	}
	piecePlacement = strings.Join(ranks, "/")

	// To Move
	toMove := "b"
	if b.ToMove {
		toMove = "w"
	}

	// Castling
	castling := "-"

	// En passant
	enPassant := "-"

	// We don't record the number of moves for now
	moves := "0 1"

	return strings.Join([]string{piecePlacement, toMove, castling, enPassant, moves}, " ")
}

func (b Board) p(rank Rank, file File) Piece {
	return b.Positions[rank][file]
}
