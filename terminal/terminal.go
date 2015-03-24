package terminal

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"sync"
)

type Drawable interface {
	Draw()
	DrawFrame()
	DrawShadow()
}

type Color termbox.Attribute

const (
	ColorDefault Color = (Color)(termbox.ColorDefault)
	ColorBlack         = (Color)(termbox.ColorBlack)
	ColorRed           = (Color)(termbox.ColorRed)
	ColorGreen         = (Color)(termbox.ColorGreen)
	ColorYellow        = (Color)(termbox.ColorYellow)
	ColorBlue          = (Color)(termbox.ColorBlue)
	ColorMagenta       = (Color)(termbox.ColorMagenta)
	ColorCyan          = (Color)(termbox.ColorCyan)
	ColorWhite         = (Color)(termbox.ColorWhite)
)

type Position struct {
	X int
	Y int
}

type Terminal struct {
	NewDrawableEvent chan Drawable
	DrawEvent        chan Drawable
	CloseEvent       chan bool
}

func New() *Terminal {

	termbox.Init()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()

	newDrawableEvent := make(chan Drawable)
	drawEvent := make(chan Drawable)
	closeEvent := make(chan bool)
	terminal := &Terminal{newDrawableEvent, drawEvent, closeEvent}

	return terminal
}

func PrintText(value interface{}, p Position) {

	text := fmt.Sprintf("%v", value)
	for i, char := range text {
		termbox.SetCell(p.X+i, p.Y, char, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func SetCell(x, y int, char rune, fg, bg Color) {
	termbox.SetCell(x, y, char, termbox.Attribute(fg), termbox.Attribute(bg))
}

func (terminal *Terminal) Handle(wg sync.WaitGroup) {

	defer wg.Done()
	defer fmt.Println("Bye bye !")
	defer termbox.Close()

	for {
		select {
		case drawable := <-terminal.NewDrawableEvent:
			drawable.DrawFrame()
		case drawable := <-terminal.DrawEvent:
			drawable.Draw()
			drawable.DrawShadow()
		case <-terminal.CloseEvent:
			return
		}

		termbox.Flush()
	}
}
