package terminal

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type terminal struct {
	CloseEvent chan bool
}

type Color termbox.Attribute

const (
	ColorDefault Color = (Color)(termbox.ColorDefault)
	ColorBlack         = (Color)(termbox.ColorBlack)
	ColorRed           = (Color)(termbox.ColorRed)
	ColorGreen         = (Color)(termbox.ColorGreen)
	ColorYellow        = (Color)(termbox.ColorYellow)
	ColorBlue          = (Color)(termbox.ColorBlue)
	ColorMagenta       = (Color)(termbox.ColorMagenta)
	ColorCyan          = (Color)(termbox.ColorCyan)
	ColorWhite         = (Color)(termbox.ColorWhite)
)

type Position struct {
	X int
	Y int
}

func NewTerminal() *terminal {

	closeEvent = make(chan bool)
	t = terminal{closeEvent}

	return &t
}

func init() {

	termbox.Init()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()
}

func (t *terminal) PrintText(value interface{}, p Position) {

	text := fmt.Sprintf("%v", value)
	for i, char := range text {
		termbox.SetCell(p.X+i, p.Y, char, termbox.ColorWhite, termbox.ColorBlack)
	}
}

/* TODO: those functions below should go to different structs */

func (t *terminal) drawBoardFrame(board Board) {

	width, height := len(board.Matrix)*2, len(board.Matrix[0])
	x, y := board.Position.X, board.Position.Y

	for i := -1; i <= width; i++ {
		ch := '-'
		if i == -1 || i == width {
			ch = '+'
		}
		termbox.SetCell(x+i, y-1, ch, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x+i, y+height, ch, termbox.ColorWhite, termbox.ColorBlack)
	}
	for i := 0; i < height; i++ {
		termbox.SetCell(x-1, y+i, '|', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x+width, y+i, '|', termbox.ColorWhite, termbox.ColorBlack)
	}
}

func (t *terminal) drawBoard(board Board) {

	for row, cells := range board.Matrix {
		for col, cell := range cells {
			x, y := board.Position.X+(row*2), board.Position.Y+col
			termbox.SetCell(x, y, '[', termbox.ColorBlack, (termbox.Attribute)(cell.Color))
			termbox.SetCell(x+1, y, ']', termbox.ColorBlack, (termbox.Attribute)(cell.Color))
		}
	}
}

func (t *terminal) drawBrickShadow(board Board) {

	bottom_frame_x := board.Position.X
	bottom_frame_y := board.Position.Y + len(board.Matrix[0])

	var border_rune rune

	for x, shadow := range board.Shadow {
		if shadow {
			border_rune = '='
		} else {
			border_rune = '-'
		}
		termbox.SetCell(bottom_frame_x+(2*x), bottom_frame_y, border_rune, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(bottom_frame_x+((2*x)+1), bottom_frame_y, border_rune, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func (terminal *terminal) Handle(boardEvent, newBoardEvent chan Board) {

	defer Wg.Done()
	defer fmt.Println("Bye bye !")
	defer termbox.Close()

	for {
		select {
		case board := <-newBoardEvent:
			terminal.drawBoardFrame(board)
		case board := <-boardEvent:
			/* TODO: this should be done in one function */
			terminal.drawBoard(board)
			terminal.drawBrickShadow(board)
		case <-terminal.closeEvent:
			return
		}

		termbox.Flush()
	}
}
