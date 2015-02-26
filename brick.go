package main

import (
	"github.com/nsf/termbox-go"
)

type Brick struct {
	Position Position
	Layout   [][]int
	Color    termbox.Attribute
	Board    *Board
	Anchored bool
}

var (
	Bricks [7]Brick
)

func init() {

	IBrick := Brick{
		Color: termbox.ColorBlue,
		Layout: [][]int{
			{1},
			{1},
			{1},
			{1}}}

	JBrick := Brick{
		Color: termbox.ColorCyan,
		Layout: [][]int{
			{1, 1, 1},
			{0, 0, 1}}}

	LBrick := Brick{
		Color: termbox.ColorYellow,
		Layout: [][]int{
			{1, 1, 1},
			{1, 0, 0}}}

	OBrick := Brick{
		Color: termbox.ColorMagenta,
		Layout: [][]int{
			{1, 1},
			{1, 1}}}

	SBrick := Brick{
		Color: termbox.ColorRed,
		Layout: [][]int{
			{0, 1, 1},
			{1, 1, 0}}}

	TBrick := Brick{
		Color: termbox.ColorWhite,
		Layout: [][]int{
			{1, 1, 1},
			{0, 1, 0}}}

	ZBrick := Brick{
		Color: termbox.ColorGreen,
		Layout: [][]int{
			{1, 1, 0},
			{0, 1, 1}}}

	Bricks[0] = IBrick
	Bricks[1] = JBrick
	Bricks[2] = LBrick
	Bricks[3] = OBrick
	Bricks[4] = SBrick
	Bricks[5] = TBrick
	Bricks[6] = ZBrick

}

func (brick *Brick) MoveLeft() {
	brick.Position.X -= 2
}

func (brick *Brick) MoveRight() {
	brick.Position.X += 2
}

func (brick *Brick) MoveDown() {
	brick.Position.Y += 1
}

func (brick *Brick) Rotate() {

	/* Transpose matrix */
	transposed := make([][]int, len(brick.Layout[0]))
	for c, _ := range transposed {
		transposed[c] = make([]int, len(brick.Layout))
	}
	for x, cells := range brick.Layout {
		for y, cell := range cells {
			transposed[y][x] = cell
		}
	}

	newLayout := make([][]int, len(brick.Layout[0]))
	for c, _ := range newLayout {
		newLayout[c] = make([]int, len(brick.Layout))
	}

	/* Change columns to rotate right */
	for x, cells := range transposed {
		for y, cell := range cells {
			newLayout[x][(len(cells)-1)-y] = cell
		}
	}

	brick.Layout = newLayout
}

func (brick *Brick) Drop() {
}
