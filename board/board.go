package board

import (
	"reflect"
	"sync"

	"github.com/grzesiek/gonetris/brick"
	"github.com/grzesiek/gonetris/terminal"
	"github.com/grzesiek/gonetris/tick"
)

type Board struct {
	Matrix              matrix
	Shadow              [10]bool
	Brick               *brick.Brick
	CloseEvent          chan bool
	BrickOperationEvent chan string
	X                   int
	Y                   int
}

func New(x, y int) *Board {

	var board Board
	board.X = x
	board.Y = y
	board.Matrix = newMatrix()
	board.CloseEvent = make(chan bool)
	board.BrickOperationEvent = make(chan string)

	return &board
}

func (board *Board) Handle(wg *sync.WaitGroup, tick *tick.Tick, terminal *terminal.Terminal) {

	defer wg.Done()
	terminal.NewDrawableEvent <- *board

	/* Create first brick */
	board.brickNext()

	for {

		select {
		case method := <-board.BrickOperationEvent:
			/* Player wants to modify brick - move, rotate, drop ... by reflection */
			reflect.ValueOf(board).MethodByName(method).Call([]reflect.Value{})
		case <-tick.TickEvent:
			board.BrickMoveDown()
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
		terminal.DrawEvent <- *board
	}
}
