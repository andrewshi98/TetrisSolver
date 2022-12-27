package tBoard

import "fmt"

type Board struct {
	w     int
	h     int
	board []int16
}

func MakeBoard(w int, h int) (*Board, error) {
	if h > 16 {
		return nil, fmt.Errorf("board does not support h > 16")
	}
	return &Board{
		w, h, make([]int16, w),
	}, nil
}
