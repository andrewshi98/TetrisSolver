package tPiece

// Rotation Available Rotations
type Rotation int

var Rotations = struct {
	L Rotation
	U Rotation
	R Rotation
	D Rotation
}{
	1, 2, 3, 4,
}

// Type Available piece types
type Type int

var Types = struct {
	I Type
	S Type
	Z Type
	L Type
	J Type
	O Type
	T Type
}{
	1, 2, 3, 4, 5, 6, 7,
}

type Piece struct {
	pType     Type
	pRotation Rotation
}
