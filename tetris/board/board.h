// board.h

#pragma once

namespace Tetris {
    class Board {
    public:
        Board(int w, int h);
        unsigned short* getBoard();
    };
} // namespace Tetris