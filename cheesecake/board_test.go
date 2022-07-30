package cheesecake

import (
	"reflect"
	"testing"
)

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want string
	}{
		{
			"test 1",
			Board{
				A1: WhiteRook, B1: WhiteKnight, C1: WhiteBishop, D1: WhiteQueen,
				E1: WhiteKing, F1: WhiteBishop, G1: WhiteKnight, H1: WhiteRook,
				A2: WhitePawn, B2: WhitePawn, C2: WhitePawn, D2: WhitePawn,
				E4: WhitePawn, F2: WhitePawn, G2: WhitePawn, H2: WhitePawn,
				A7: BlackPawn, B7: BlackPawn, C7: BlackPawn, D7: BlackPawn,
				E5: BlackPawn, F7: BlackPawn, G7: BlackPawn, H7: BlackPawn,
				A8: BlackRook, B8: BlackKnight, C8: BlackBishop, D8: BlackQueen,
				E8: BlackKing, F8: BlackBishop, G8: BlackKnight, H8: BlackRook,
			},
			`  ╔═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╗
8 ║ r │ n │ b │ q │ k │ b │ n │ r ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
7 ║ p │ p │ p │ p │   │ p │ p │ p ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
6 ║   │   │   │   │   │   │   │   ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
5 ║   │   │   │   │ p │   │   │   ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
4 ║   │   │   │   │ P │   │   │   ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
3 ║   │   │   │   │   │   │   │   ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
2 ║ P │ P │ P │ P │   │ P │ P │ P ║
  ╟───┼───┼───┼───┼───┼───┼───┼───╢
1 ║ R │ N │ B │ Q │ K │ B │ N │ R ║
  ╚═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╝
    a   b   c   d   e   f   g   h  
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Board.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Place(t *testing.T) {
	tests := []struct {
		name    string
		b       Board
		p       Piece
		s       Square
		want    Board
		wantErr bool
	}{
		{"test 1", Board{}, WhiteBishop, E4, Board{E4: WhiteBishop}, false},
		{"test 2", Board{E4: WhiteBishop}, BlackRook, E7, Board{E4: WhiteBishop, E7: BlackRook}, false},
		{"test 3: invalid piece", Board{}, 50, E7, Board{}, true},
		{"test 4: invalid square", Board{}, BlackRook, 100, Board{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.b
			p := tt.p
			s := tt.s

			err := b.Place(p, s)
			got := b
			want := tt.want

			if (err != nil) != tt.wantErr {
				t.Errorf("Board.Place() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(b, want) {
				t.Errorf("Board.Place() =\n%v\nwant\n%v", got, want)
			}
		})
	}
}

func TestBoard_Clear(t *testing.T) {
	tests := []struct {
		name    string
		b       Board
		s       Square
		want    Board
		wantErr bool
	}{
		{
			"test 1: empty board, valid square",
			Board{},
			E4,
			Board{},
			false,
		},
		{
			"test 2: ",
			Board{E4: BlackRook},
			E4,
			Board{},
			false,
		},
		{
			"test 2: invalid square",
			Board{E4: BlackRook},
			100,
			Board{E4: BlackRook},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.b
			s := tt.s

			err := b.Clear(s)

			got := b
			want := tt.want

			if (err != nil) != tt.wantErr {
				t.Errorf("Board.Clear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(b, want) {
				t.Errorf("Board.Clear() =\n%v\nwant\n%v", got, want)
			}
		})
	}
}
