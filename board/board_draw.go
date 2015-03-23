package board

import (
	"github.com/grzesiek/gonetris/terminal"
)

func (board board) DrawFrame() {

	width, height := len(board.Matrix)*2, len(board.Matrix[0])
	x, y := board.X, board.Y

	for i := -1; i <= width; i++ {
		ch := '-'
		if i == -1 || i == width {
			ch = '+'
		}
		terminal.SetCell(x+i, y-1, ch, terminal.ColorWhite, terminal.ColorBlack)
		terminal.SetCell(x+i, y+height, ch, terminal.ColorWhite, terminal.ColorBlack)
	}
	for i := 0; i < height; i++ {
		terminal.SetCell(x-1, y+i, '|', terminal.ColorWhite, terminal.ColorBlack)
		terminal.SetCell(x+width, y+i, '|', terminal.ColorWhite, terminal.ColorBlack)
	}
}

func (board board) Draw() {

	for row, cells := range board.Matrix {
		for col, cell := range cells {
			x, y := board.X+(row*2), board.Y+col
			terminal.SetCell(x, y, '[', terminal.ColorBlack, cell.Color)
			terminal.SetCell(x+1, y, ']', terminal.ColorBlack, cell.Color)
		}
	}
}

func (board board) DrawShadow() {

	bottom_frame_x := board.X
	bottom_frame_y := board.Y + len(board.Matrix[0])

	var border_rune rune

	for x, shadow := range board.Shadow {
		if shadow {
			border_rune = '='
		} else {
			border_rune = '-'
		}
		terminal.SetCell(bottom_frame_x+(2*x), bottom_frame_y, border_rune, terminal.ColorWhite, terminal.ColorBlack)
		terminal.SetCell(bottom_frame_x+((2*x)+1), bottom_frame_y, border_rune, terminal.ColorWhite, terminal.ColorBlack)
	}
}
