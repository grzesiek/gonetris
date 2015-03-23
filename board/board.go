package board

import (
	"reflect"
	"sync"

	"github.com/grzesiek/gonetris/brick"
	"github.com/grzesiek/gonetris/terminal"
)

type position struct {
	X, Y int
}
type board struct {
	Matrix         matrix
	Shadow         [10]bool
	Brick          *brick.Brick
	DrawEvent      chan terminal.Drawable
	CloseEvent     chan bool
	BrickOperation chan string
	X              int
	Y              int
}

func New(x, y int) *board {

	var board board
	board.X = x
	board.Y = y
	board.Matrix = newMatrix()
	board.DrawEvent = make(chan terminal.Drawable)
	board.CloseEvent = make(chan bool)

	/* TODO: this should go to Brick or be changed in other way
	boardBrickOperation = make(chan string)
	*/

	/* TODO: This shoudln't be needed anymore
	TerminalNewboardEvent <- board
	*/

	return &board
}

func (board *board) Handle(wg sync.WaitGroup) {

	defer wg.Done()

	/* Create first brick */
	board.brickNext()

	for {

		select {
		case method := <-board.BrickOperation:
			/* Player wants to modify brick - move, rotate, drop ... by reflection */
			/* This also handles moving down bick on tick */
			reflect.ValueOf(board).MethodByName(method).Call([]reflect.Value{})
		case <-board.CloseEvent:
			return
		}

		/* Reset empty cells (not filled) */
		board.Matrix.resetEmptyCells()
		/* Draw current brick board */
		board.brickDraw()
		/* Set current brick shadow */
		board.brickSetShadow()

		/* User can move birck one last time after it touches something */
		if board.needsNextBrick() {

			/* Fill with current brick*/
			board.fillWithBrick()
			/* Chose next brick */
			board.brickNext()
			/* Remove full lines */
			board.Matrix.removeFullLines()
		}

		/* emit boardEvent */
		board.DrawEvent <- *board
	}
}
