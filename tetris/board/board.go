package tBoard

import (
	"fmt"
	"github.com/andrewshi98/TetrisSolver/tetris/config"
)

type Board struct {
	w     int
	h     int
	board []config.YData
}

func MakeBoard(w int, h int) (*Board, error) {
	if h > 16 {
		return nil, fmt.Errorf("board does not support h > 16")
	}
	return &Board{
		w, h, make([]int16, w),
	}, nil
}
