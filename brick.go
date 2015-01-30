package main

import (
	"github.com/nsf/termbox-go"
)

type Brick struct {
	Position Position
	Layout   [][]int
	Color    termbox.Attribute
	Board    *Board
}

type Event uint16

const (
	BrickMoveDown Event = iota
	BrickMoveLeft
	BrickMoveRight
	BrickRotate
)

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

func (b *Brick) MoveLeft() {
	if !b.Board.BrickTouched(BorderLeft, true) {
		b.Position.X -= 2
	}
}

func (b *Brick) MoveRight() {
	if !b.Board.BrickTouched(BorderRight, true) {
		b.Position.X += 2
	}
}

func (b *Brick) MoveDown() {
	/* Check if bricked touch something */
	if b.Board.BrickTouched(BorderBottom, false) {
		/* Fill with current brick*/
		b.Board.FillWithBrick()
		/* Chose next brick */
		b.Board.NextBrick()
	} else {
		b.Position.Y += 1
	}
}

func (b *Brick) Rotate() {

	if !b.Board.BrickTouched(BorderLeft|BorderRight, true) {
		newLayout := make([][]int, len(b.Layout[0]))
		for c, _ := range newLayout {
			newLayout[c] = make([]int, len(b.Layout))
		}
		for x, cells := range b.Layout {
			for y, cell := range cells {
				newLayout[y][x] = cell
			}
		}
		b.Layout = newLayout
	}
}

func (b *Brick) Drop() {
}
