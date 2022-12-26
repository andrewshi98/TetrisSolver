#include <vector>

#include <tetris/board/board.h>
#include <tetris/piece/piece.h>

#pragma once


namespace Tetris
{
    bool piece_reachable(Tetris::Board board, Tetris::Piece piece, int x, int y);

    // Low Level Op
    bool place_piece(Tetris::Board board, Tetris::Placing placing);
    bool unplace_piece(Tetris::Board board, Tetris::Placing placing);
    std::vector<Tetris::Placing> get_placing_permutations(Tetris::Board board, Tetris::Piece piece);

    // High Level Op
    // TODO: Come up with it...

    // data structures and types
    struct Placing {
        Tetris::Piece piece;
        int x;
        int y;
    };

} // namespace Game
