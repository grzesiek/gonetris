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

type BrickEvent uint16

const (
	BrickMoveDown BrickEvent = iota
	BrickNew
	BrickMoveLeft
	BrickMoveRight
)

var (
	Bricks         [7]Brick
	BricksChan     = make(chan *Brick)
	BrickEventChan = make(chan BrickEvent)
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

func NextBrick() *Brick {
	rand.Seed(time.Now().UTC().UnixNano())
	brick := &Bricks[rand.Intn(7)]
	brick.Position = Position{0, 0}
	return brick
}

func HandleBricks() {

	defer Wg.Done()
	brick := NextBrick()
	BricksChan <- brick

	for e := range BrickEventChan {

		switch e {
		case BrickNew:
			/* Change current brick */
			brick = NextBrick()
			BricksChan <- brick
		case BrickMoveDown:
			/* Lower brick */
			brick.MoveDown()
		}

	}
}
