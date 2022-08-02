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

func TestSquare_File(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want File
	}{
		{"A1", A1, FileA},
		{"A8", A8, FileA},
		{"H1", H1, FileH},
		{"H8", H8, FileH},
		{"99", 99, NoFile},
		{"NoSquare", NoSquare, NoFile},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.File(); got != tt.want {
				t.Errorf("Square.File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Rank(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want Rank
	}{
		{"A1", A1, Rank1},
		{"A8", A8, Rank8},
		{"H1", H1, Rank1},
		{"H8", H8, Rank8},
		{"99", 99, NoRank},
		{"NoSquare", NoSquare, NoRank},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Rank(); got != tt.want {
				t.Errorf("Square.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeSquare(t *testing.T) {
	tests := []struct {
		name string
		f    File
		r    Rank
		want Square
	}{
		{"FileA Rank1", FileA, Rank1, A1},
		{"FileA Rank8", FileA, Rank8, A8},
		{"FileH Rank1", FileH, Rank1, H1},
		{"FileH Rank8", FileH, Rank8, H8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeSquare(tt.f, tt.r); got != tt.want {
				t.Errorf("NewSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnSameDiag(t *testing.T) {
	tests := []struct {
		name string
		s1   Square
		s2   Square
		want bool
	}{
		{"test 1", D4, G7, true},
		{"test 2", D4, B6, true},
		{"test 3", D4, F2, true},
		{"test 4", D4, G1, true},
		{"test 5", D4, D6, false},
		{"test 6", D4, G4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnSameDiag(tt.s1, tt.s2); got != tt.want {
				t.Errorf("OnSameDiag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHaveSameColor(t *testing.T) {
	tests := []struct {
		name string
		p1   Piece
		p2   Piece
		want bool
	}{
		{"test 1", BlackRook, BlackPawn, true},
		{"test 2", BlackRook, WhiteRook, false},
		{"test 3", BlackRook, NoPiece, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HaveSameColor(tt.p1, tt.p2); got != tt.want {
				t.Errorf("HaveSameColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Right(t *testing.T) {
	tests := []struct {
		name string
		f    File
		want File
	}{
		{"test 1", NoFile, NoFile},
		{"test 2", 9, NoFile},
		{"test 3", -99, NoFile},
		{"test 4", FileA, FileB},
		{"test 5", FileH, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Right(); got != tt.want {
				t.Errorf("File.Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Left(t *testing.T) {
	tests := []struct {
		name string
		f    File
		want File
	}{
		{"test 1", NoFile, NoFile},
		{"test 2", 9, NoFile},
		{"test 3", -99, NoFile},
		{"test 4", FileH, FileG},
		{"test 5", FileA, NoFile},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Left(); got != tt.want {
				t.Errorf("File.Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_Up(t *testing.T) {
	tests := []struct {
		name string
		r    Rank
		want Rank
	}{
		{"test 1", NoRank, NoRank},
		{"test 2", 9, NoRank},
		{"test 3", -99, NoRank},
		{"test 4", Rank1, Rank2},
		{"test 5", Rank8, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Up(); got != tt.want {
				t.Errorf("Rank.Up() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_Down(t *testing.T) {
	tests := []struct {
		name string
		r    Rank
		want Rank
	}{
		{"test 1", NoRank, NoRank},
		{"test 2", 9, NoRank},
		{"test 3", -99, NoRank},
		{"test 4", Rank1, NoRank},
		{"test 5", Rank8, Rank7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Down(); got != tt.want {
				t.Errorf("Rank.Down() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_UpRight(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want Square
	}{
		{"test 1", E4, F5},
		{"test 2", H1, NoSquare},
		{"test 3", H8, NoSquare},
		{"test 4", NoSquare, NoSquare},
		{"test 5", 100, NoSquare},
		{"test 6", -1, NoSquare},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpRight(); got != tt.want {
				t.Errorf("Square.UpRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_UpLeft(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want Square
	}{
		{"test 1", E4, D5},
		{"test 2", A1, NoSquare},
		{"test 3", A8, NoSquare},
		{"test 4", NoSquare, NoSquare},
		{"test 5", 100, NoSquare},
		{"test 6", -1, NoSquare},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpLeft(); got != tt.want {
				t.Errorf("Square.UpLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_DownRight(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want Square
	}{
		{"test 1", E4, F3},
		{"test 2", A1, NoSquare},
		{"test 3", H1, NoSquare},
		{"test 4", NoSquare, NoSquare},
		{"test 5", 100, NoSquare},
		{"test 6", -1, NoSquare},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DownRight(); got != tt.want {
				t.Errorf("Square.DownRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_DownLeft(t *testing.T) {
	tests := []struct {
		name string
		s    Square
		want Square
	}{
		{"test 1", E4, D3},
		{"test 2", A8, NoSquare},
		{"test 3", A1, NoSquare},
		{"test 4", H1, NoSquare},
		{"test 5", NoSquare, NoSquare},
		{"test 6", 100, NoSquare},
		{"test 7", -1, NoSquare},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DownLeft(); got != tt.want {
				t.Errorf("Square.DownLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_BishopCanMove(t *testing.T) {
	tests := []struct {
		name   string
		b      Board
		origin Square
		target Square
		want   bool
	}{
		{
			"test 1: same color target",
			Board{D4: WhiteBishop, G7: WhiteRook},
			D4,
			G7,
			false,
		},
		{
			"test 2: piece in between",
			Board{D4: WhiteBishop, F6: WhiteKnight, G7: BlackRook},
			D4,
			G7,
			false,
		},
		{
			"test 3: valid move",
			Board{D4: WhiteBishop, G7: BlackRook},
			D4,
			G7,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.BishopCanMove(tt.origin, tt.target); got != tt.want {
				t.Errorf("Board.BishopCanMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
