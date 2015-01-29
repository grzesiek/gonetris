package main

import (
	"github.com/nsf/termbox-go"
)

type Brick struct {
	Position Position
	Layout   [][]int
	Color    termbox.Attribute
}

type Event uint16

const (
	BrickMoveDown Event = iota
	BrickMoveLeft
	BrickMoveRight
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
}

func (b *Brick) MoveRight() {
}

func (b *Brick) MoveDown() {
	b.Position.Y += 1
}

func (b *Brick) Drop() {
}
