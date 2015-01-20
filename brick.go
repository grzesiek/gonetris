package main

import ()

type Brick interface {
	Draw()
	MoveLeft()
	MoveRight()
	Drop()
}

var (
	CurrentBrick *Brick
	Bricks       [7]Brick
)

func init() {
}
