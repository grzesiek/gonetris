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
			if cell == 1 && y > -1 {
				board.Matrix[x][y].Color = brick.Color
				board.Matrix[x][y].Empty = false
			}
		}
	}

}

func (board *Board) BrickShadowDraw() {

	brick := board.Brick
	layout := brick.Layout
	minX := len(layout)
	maxX := 0

	for x, cells := range layout {
		for _, cell := range cells {
			if cell == 1 {
				if x < minX {
					minX = x
				}

				if x > maxX {
					maxX = x
				}
			}
		}
	}

	minX += brick.Position.X
	maxX += brick.Position.X

	for x := range board.Shadow {
		board.Shadow[x] = x >= minX && x <= maxX
	}
}

func (board *Board) brickTouched(blocker BrickBlocker) bool {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+bx, brick.Position.Y+by
			if cell == 1 && y > -1 {

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
			if cell == 1 && y > -1 {
				/* Check if x index > matrix capacity */
				if x > len(board.Matrix)-1 {
					return false
				}

				/* Check if x index < matrix capacity */
				if x < 0 {
					return false
				}

				/* Check if y index > matrix capacity */
				if y > len(board.Matrix[0])-1 {
					return false
				}

				/* Check if there is already embedded brick */
				if board.Matrix[x][y].Embedded {
					return false /* TODO: rotation bug */
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
			if cell == 1 && y > -1 {
				board.Matrix[x][y].Embedded = true
			}
		}
	}
}

func (board *Board) RemoveFullLines() int {

	var lineFull bool
	var removedLines []int

	for by := 0; by < len(board.Matrix[0]); by++ {
		lineFull = true
		for bx := 0; bx < len(board.Matrix); bx++ {
			lineFull = lineFull && board.Matrix[bx][by].Embedded
		}

		if lineFull {
			for bx := 0; bx < len(board.Matrix); bx++ {
				board.ResetCell(bx, by)
			}
			removedLines = append(removedLines, by)
		}
	}

	if len(removedLines) > 0 {
		for _, y := range removedLines {
			for by := y - 1; by > 0; by-- {
				for bx := 0; bx < len(board.Matrix); bx++ {
					board.Matrix[bx][by+1] = board.Matrix[bx][by]
					board.ResetCell(bx, by)
				}
			}
		}
	}

	return len(removedLines)
}

func (board *Board) BrickNext() *Brick {
	rand.Seed(time.Now().UTC().UnixNano())
	brick := &Bricks[rand.Intn(7)]
	brick.Position = Position{4, brick.StartOffset - 1}
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
