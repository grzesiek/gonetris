package main

import (
	"math/rand"
	"time"
)

func (board *Board) BrickDraw() {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+bx, brick.Position.Y+by
			if cell == 1 {
				board.Matrix[x][y].Color = brick.Color
				board.Matrix[x][y].Empty = false
			}
		}
	}

}

func (board *Board) brickTouched(blocker BrickBlocker) bool {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+bx, brick.Position.Y+by
			if cell == 1 {

				if blocker&BorderRight != 0 {
					/* Touched right border */
					if len(board.Matrix) == x+1 {
						return true
					}
				}
				if blocker&BorderLeft != 0 {
					/* Touched left border */
					if x == 0 {
						return true
					}
				}
				if blocker&BorderBottom != 0 {
					/* Touched bottom border */
					if len(board.Matrix[0]) == y+1 {
						return true
					}
				}
				if blocker&BrickBelow != 0 {
					/* Touched other brick, that already filled board at the bottom */
					if y+1 < len(board.Matrix[0]) && board.Matrix[x][y+1].Embedded {
						return true
					}
				}
				/* Check below conditions only if we are moving horizontally */
				if blocker&BrickAtLeft != 0 {
					/* Touched other brick, that already filled board at left */
					if x > 1 && board.Matrix[x-1][y].Embedded {
						return true
					}
				}
				if blocker&BrickAtRight != 0 {
					/* Touched other brick, that already filled board at right */
					if x+1 < len(board.Matrix) && board.Matrix[x+1][y].Embedded {
						return true
					}
				}

			}
		}
	}

	return false
}

func (board *Board) brickCanRotate() bool {

	if !board.brickTouched(Something) {
		return true
	}

	brick := board.Brick
	rotationPredictionLayout := brick.RotationLayout()

	for bx, cells := range rotationPredictionLayout {
		for by, cell := range cells {
			x, y := brick.Position.X+bx, brick.Position.Y+by
			if cell == 1 {
				/* Check if x index > matrix */
				if x > len(board.Matrix)-1 {
					return false
				}

				/* Check if x index < matrix */
				if x < 0 {
					return false
				}

				/* Check if y index > matrix */
				if y > len(board.Matrix[0])-1 {
					return false
				}

				/* Embedded */
				if board.Matrix[x][y].Embedded {
					return false
				}

			}
		}
	}

	return true
}

func (board *Board) FillWithBrick() {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+bx, brick.Position.Y+by
			if cell == 1 {
				board.Matrix[x][y].Embedded = true
			}
		}
	}
}

func (board *Board) BrickNext() *Brick {
	rand.Seed(time.Now().UTC().UnixNano())
	brick := &Bricks[rand.Intn(7)]
	brick.Position = Position{0, 0}
	brick.Anchored = false
	board.Brick = brick
	brick.Board = board

	return brick
}

func (board *Board) BrickMoveLeft() {

	if !board.brickTouched(BorderLeft | BrickAtLeft) {
		board.Brick.MoveLeft()
	}
}

func (board *Board) BrickMoveRight() {

	if !board.brickTouched(BorderRight | BrickAtRight) {
		board.Brick.MoveRight()
	}
}

func (board *Board) BrickMoveDown() {

	if !board.brickTouched(BorderBottom | BrickBelow) {
		board.Brick.MoveDown()
	}
}

func (board *Board) BrickRotate() {

	if board.brickCanRotate() {
		board.Brick.Rotate()
	}
}

func (board *Board) BrickDrop() {

	for !board.brickTouched(BorderBottom | BrickBelow) {
		board.BrickMoveDown()
	}
}

func (board *Board) NeedsNextBrick() bool {

	/* Brick becomes anchored once it touches something below at the first time */
	/* User can move birck one last time after it touches something */

	touched := board.brickTouched(BorderBottom | BrickBelow)
	anchored := board.Brick.Anchored
	if touched {
		board.Brick.Anchored = true
	}
	return touched && anchored
}
