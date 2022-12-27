package tetris

import (
	tBoard "github.com/andrewshi98/TetrisSolver/tetris/board"
	tPiece "github.com/andrewshi98/TetrisSolver/tetris/piece"
)

// data structures and types
type Placing struct {
	piece tPiece.Piece
	x     int
	y     int
}

type Tetris interface {
	// Basic Operations
	PieceReachable(board tBoard.Board, piece tPiece.Piece, x int, y int) bool

	// Low Level Op
	PlacePiece(board tBoard.Board, placing Placing) bool
	UnPlacePiece(board tBoard.Board, placing Placing) bool
	GetPlacingPermutations(board tBoard.Board, piece tPiece.Piece) []Placing

	// High Level Op
	// TODO: Come up with it...

} // namespace Game
