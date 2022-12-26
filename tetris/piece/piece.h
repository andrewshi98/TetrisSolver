#pragma once

// Piece with type and rotation data
namespace Tetris {

    // Available Piece types
    enum PieceType {
        I, S, Z, L, J, O, T
    };

    // Available Rotations (Pointless to use enum but...)
    enum PieceRotation {
        L, U, R, D
    };

    struct Piece {
        PieceType type;
        PieceRotation rotation;
    };
} // namespace Tetris