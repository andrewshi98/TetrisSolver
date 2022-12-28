package tPiece

import "github.com/andrewshi98/TetrisSolver/tetris/config"

// PieceData Always Square. We can derive the width and length from that.
type PieceData []config.YData
type RotationData [4]PieceData

// Named piece rotation data
var (
	IRotations = RotationData{
		{4, 4, 4, 4},
		{0, 0, 15, 0},
		{2, 2, 2, 2},
		{0, 15, 0, 0},
	}
	JRotations = RotationData{
		{6, 2, 2},
		{0, 7, 4},
		{2, 2, 3},
		{1, 7, 0},
	}
	LRotations = RotationData{
		{2, 2, 6},
		{0, 7, 1},
		{3, 2, 2},
		{4, 7, 0},
	}
	ORotations = RotationData{
		{3, 3},
		{3, 3},
		{3, 3},
		{3, 3},
	}
	SRotations = RotationData{
		{2, 6, 4},
		{0, 6, 3},
		{1, 3, 2},
		{6, 3, 0},
	}
	TRotations = RotationData{
		{2, 6, 2},
		{0, 7, 2},
		{2, 3, 2},
		{2, 7, 0},
	}
	ZRotations = RotationData{
		{4, 6, 2},
		{0, 3, 6},
		{2, 3, 1},
		{3, 6, 0},
	}
)

// a
var (
	Type2RotationD = 
)
