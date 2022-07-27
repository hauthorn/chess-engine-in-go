package own_chess

import (
	"reflect"
	"testing"
)

func TestBoard_getPieceAt(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		fields Board
		args   args
		want   *Piece
	}{
		{"Black rook A8", DefaultBoard(), args{position: Position{A, 8}}, &BR},
		{"White Queen D1", DefaultBoard(), args{position: Position{D, 1}}, &WQ},
		{"White Pawn E2", DefaultBoard(), args{position: Position{E, 2}}, &WP},
		{"Empty E4", DefaultBoard(), args{position: Position{E, 4}}, &EM},
		{"Black rook H8", DefaultBoard(), args{position: Position{H, 8}}, &BR},
		{"White rook A1", DefaultBoard(), args{position: Position{A, 1}}, &WR},
		{"Nonstandard: White king D2", KingsAndRook(WHITE), args{position: Position{D, 2}}, &WK},
		{"Nonstandard: White rook E2", KingsAndRook(WHITE), args{position: Position{E, 2}}, &WR},
		{"Nonstandard: Black king D8", KingsAndRook(WHITE), args{position: Position{D, 8}}, &BK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Board{
				Positions: tt.fields.Positions,
				ToMove:    tt.fields.ToMove,
			}
			if got := b.getPieceAt(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPieceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Clone(t *testing.T) {
	clone := DefaultBoard().Clone()

	tests := []struct {
		name     string
		position Position
		piece    *Piece
	}{
		{"White Rook A1", Position{A, 1}, &WR},
		{"White Rook H1", Position{H, 1}, &WR},
		{"White Queen D1", Position{D, 1}, &WQ},
		{"Black Queen D8", Position{D, 8}, &BQ},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clone.getPieceAt(tt.position); !reflect.DeepEqual(got, tt.piece) {
				t.Errorf("Position %v was %v, want %v", tt.position, got, tt.piece)
			}
		})
	}
}

func TestBoard_FENString(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  string
	}{
		{
			"Default board FEN string",
			DefaultBoard(),
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
		},
		{
			"KingsAndRook",
			KingsAndRook(true),
			"3k4/8/8/8/8/8/3KR3/8 w - - 0 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.board
			if got := b.FENString(); got != tt.want {
				t.Errorf("FENString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPiece_String(t *testing.T) {
	type fields struct {
		Color Color
		Type  PieceType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"White king", fields{WHITE, King}, "K"},
		{"Black king", fields{BLACK, King}, "k"},
		{"Black rook", fields{BLACK, Rook}, "r"},
		{"White queen", fields{WHITE, Queen}, "Q"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Piece{
				Color: tt.fields.Color,
				Type:  tt.fields.Type,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
