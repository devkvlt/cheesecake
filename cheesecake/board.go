package cheesecake

import (
	"errors"
	"fmt"
	"strings"
)

// Piece represents a chess piece (type and color.)
type Piece int

// Square represents a square on a chess board (file and rank.)
type Square int

// Board represents a chess board with pieces in it.
type Board map[Square]Piece

// 12 chess pieces and NoPiece which represents an invalid piece.
const (
	NoPiece Piece = iota
	WhiteKing
	WhiteQueen
	WhiteRook
	WhiteBishop
	WhiteKnight
	WhitePawn
	BlackKing
	BlackQueen
	BlackRook
	BlackBishop
	BlackKnight
	BlackPawn
)

// 64 aquares of a chess board and NoSquare which represents an invalid square.
const (
	NoSquare Square = iota
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// DefaultBoard represents a board in the initial position.
var DefaultBoard = Board{
	A1: WhiteRook, B1: WhiteKnight, C1: WhiteBishop, D1: WhiteQueen,
	E1: WhiteKing, F1: WhiteBishop, G1: WhiteKnight, H1: WhiteRook,
	A2: WhitePawn, B2: WhitePawn, C2: WhitePawn, D2: WhitePawn,
	E2: WhitePawn, F2: WhitePawn, G2: WhitePawn, H2: WhitePawn,
	A7: BlackPawn, B7: BlackPawn, C7: BlackPawn, D7: BlackPawn,
	E7: BlackPawn, F7: BlackPawn, G7: BlackPawn, H7: BlackPawn,
	A8: BlackRook, B8: BlackKnight, C8: BlackBishop, D8: BlackQueen,
	E8: BlackKing, F8: BlackBishop, G8: BlackKnight, H8: BlackRook,
}

// String implements the fmt.Stringer interface.
func (p Piece) String() string {
	return map[Piece]string{
		WhiteRook:   "R",
		WhiteKnight: "N",
		WhiteBishop: "B",
		WhiteQueen:  "Q",
		WhiteKing:   "K",
		WhitePawn:   "P",
		BlackRook:   "r",
		BlackKnight: "n",
		BlackBishop: "b",
		BlackQueen:  "q",
		BlackKing:   "k",
		BlackPawn:   "p",
		NoPiece:     " ",
	}[p]
}

// IsValid reports whether a piece is one of the 12 valid pieces.
func (p Piece) IsValid() bool {
	return p >= 1 && p <= 12
}

// IsValid reports whether a square is one of the 64 valid squares.
func (s Square) IsValid() bool {
	return s >= 1 && s <= 64
}

// MakeBoard returns an empty board.
func MakeBoard() Board {
	return Board{}
}

// MakeDefaultBoard returns a board in the initial position.
func MakeDefaultBoard() Board {
	return DefaultBoard
}

// String implements the fmt.Stringer interface.
func (b Board) String() string {
	top := "  ╔═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╗\n"
	sep := "  ╟───┼───┼───┼───┼───┼───┼───┼───╢\n"
	bot := "  ╚═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╝\n    a   b   c   d   e   f   g   h  \n"
	ranks := []string{
		fmt.Sprintf("8 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A8], b[B8], b[C8], b[D8], b[E8], b[F8], b[G8], b[H8]),
		fmt.Sprintf("7 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A7], b[B7], b[C7], b[D7], b[E7], b[F7], b[G7], b[H7]),
		fmt.Sprintf("6 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A6], b[B6], b[C6], b[D6], b[E6], b[F6], b[G6], b[H6]),
		fmt.Sprintf("5 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A5], b[B5], b[C5], b[D5], b[E5], b[F5], b[G5], b[H5]),
		fmt.Sprintf("4 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A4], b[B4], b[C4], b[D4], b[E4], b[F4], b[G4], b[H4]),
		fmt.Sprintf("3 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A3], b[B3], b[C3], b[D3], b[E3], b[F3], b[G3], b[H3]),
		fmt.Sprintf("2 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A2], b[B2], b[C2], b[D2], b[E2], b[F2], b[G2], b[H2]),
		fmt.Sprintf("1 ║ %s │ %s │ %s │ %s │ %s │ %s │ %s │ %s ║\n", b[A1], b[B1], b[C1], b[D1], b[E1], b[F1], b[G1], b[H1]),
	}

	return top + strings.Join(ranks, sep) + bot
}

// Place puts the piece p on the square s. If p or s is invalid, an non-nil
// error is returned.
func (b Board) Place(p Piece, s Square) error {
	if !p.IsValid() {
		return errors.New("invalid piece")
	}

	if !s.IsValid() {
		return errors.New("invalid square")
	}

	b[s] = p
	return nil
}

// Clear removes the piece on square s. If s is empty, Clear is a no-op. If s is
// invalid, a non-nil error is returned.
func (b Board) Clear(s Square) error {
	if !s.IsValid() {
		return errors.New("invalid square")
	}

	delete(b, s)
	return nil
}

// Move moves the piece that is on the origin square to the target square. If
// origin or target is invalid or there is no piece on origin, a non-nil error
// is returned.
func (b Board) Move(origin, target Square) error {
	if !origin.IsValid() {
		return errors.New("invalid origin square")
	}
	if !target.IsValid() {
		return errors.New("invalid target square")
	}

	p, ok := b[origin]
	if !ok {
		return errors.New("origin square is empty")
	}
	if !p.IsValid() {
		return errors.New("invalid piece on origin square")
	}

	b[target] = p
	delete(b, origin)
	return nil
}
