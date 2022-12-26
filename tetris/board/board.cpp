#include <iostream>
#include <vector>
#include "board.h"

// Data Declearation

using namespace Tetris;

class Board {
private:
    // Data alignment: (x, y)
    //  Y bit first, ordered by Columns
    int w;
    int h;
    unsigned short board[];
public:
    Board(int width, int height) {
        if (height > 16) {
            throw std::invalid_argument("board do not support height > 16");
        }
        w = width;
        h = height;
    }

    unsigned short* getBoard() {
        return board;
    }
};

int main() {
    int a = 5;
    
    std::cout << "hi" << std::endl;
}