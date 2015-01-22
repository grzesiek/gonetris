package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Brick struct {
	Position Position
	Layout   [][]int
	Color    termbox.Attribute
}

var (
	CurrentBrick *Brick
	Bricks       [7]Brick
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

func (b *Brick) DrawOnBoard() {

	for bx, cells := range b.Layout {
		for by, cell := range cells {
			x, y := b.Position.X+(bx*2), b.Position.Y+by
			if cell == 1 {
				MyPlayer.Board.Matrix[x][y].Char.Ch = '['
				MyPlayer.Board.Matrix[x+1][y].Char.Ch = ']'
				MyPlayer.Board.Matrix[x][y].Char.Bg = b.Color
				MyPlayer.Board.Matrix[x+1][y].Char.Bg = b.Color
				MyPlayer.Board.Matrix[x][y].Char.Fg = termbox.ColorBlack
				MyPlayer.Board.Matrix[x+1][y].Char.Fg = termbox.ColorBlack
			}
		}
	}

}

func (b *Brick) MoveLeft() {
}

func (b *Brick) MoveRight() {
}

func (b *Brick) MoveDown() {
	b.Position.Y -= 1
}

func (b *Brick) Drop() {
}

func NextBrick() {
	rand.Seed(time.Now().UTC().UnixNano())
	CurrentBrick = &Bricks[rand.Intn(7)]
	CurrentBrick.Position = Position{0, 0}
}
