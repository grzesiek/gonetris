package brick

import "github.com/grzesiek/gonetris/terminal"

type Brick struct {
	Layout      [][]int
	Color       terminal.Color
	Position    terminal.Position
	Anchored    bool
	StartOffset int
}

var (
	Bricks [7]Brick
)

func init() {

	IBrick := Brick{
		Color:       ColorBlue,
		StartOffset: -1,
		Layout: [][]int{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0}}}

	JBrick := Brick{
		Color:       ColorCyan,
		StartOffset: -2,
		Layout: [][]int{
			{0, 1, 1},
			{0, 1, 0},
			{0, 1, 0}}}

	LBrick := Brick{
		Color:       ColorYellow,
		StartOffset: -2,
		Layout: [][]int{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 1}}}

	OBrick := Brick{
		Color:       ColorMagenta,
		StartOffset: -1,
		Layout: [][]int{
			{1, 1},
			{1, 1}}}

	SBrick := Brick{
		Color:       ColorRed,
		StartOffset: -1,
		Layout: [][]int{
			{1, 0},
			{1, 1},
			{0, 1}}}

	TBrick := Brick{
		Color:       ColorWhite,
		StartOffset: -2,
		Layout: [][]int{
			{0, 1, 0},
			{0, 1, 1},
			{0, 1, 0}}}

	ZBrick := Brick{
		Color:       ColorGreen,
		StartOffset: -1,
		Layout: [][]int{
			{0, 1},
			{1, 1},
			{1, 0}}}

	Bricks[0] = IBrick
	Bricks[1] = JBrick
	Bricks[2] = LBrick
	Bricks[3] = OBrick
	Bricks[4] = SBrick
	Bricks[5] = TBrick
	Bricks[6] = ZBrick

}

func (brick *Brick) MoveLeft() {
	brick.Position.X -= 1
}

func (brick *Brick) MoveRight() {
	brick.Position.X += 1
}

func (brick *Brick) MoveDown() {
	brick.Position.Y += 1
}

func (brick *Brick) RotationLayout() [][]int {

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

	return newLayout
}

func (brick *Brick) Rotate() {

	brick.Layout = brick.RotationLayout()
}
