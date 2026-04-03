package mouse

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MouseForm interface {
	Collison() bool
}

type ButtonEllipse struct {
	X     int
	Y     int
	XAxis int
	YAxis int
}

func (E *ButtonEllipse) Collission() bool {
	x, y := ebiten.CursorPosition()

	if E.ellipse_formula(x, y) {
		return true
	}
	return false
}

func (E *ButtonEllipse) ellipse_formula(x int, y int) bool {
	formula := ((x-E.X)^2)/(E.XAxis^2) + ((y-E.Y)^2)/(E.YAxis^2)
	return formula >= 1
}

type ButtonRound struct {
	X      int
	Y      int
	Raduis int
}

func (R *ButtonRound) Collission() bool {
	x, y := ebiten.CursorPosition()
	if (x <= R.X+R.Raduis && x >= R.X-R.Raduis) && (y <= R.Y+R.Raduis && y >= R.Y-R.Raduis) {
		return true
	}
	return false
}

type ButtonRectangle struct {
	X      int // upper left
	Y      int //
	Lenght int
	Width  int
}

func (B *ButtonRectangle) Collission() bool {
	x, y := ebiten.CursorPosition()
	if (x <= B.X) && (x >= B.X+B.Lenght) && (y <= B.Y) && (y >= B.Y-B.Width) {
		return true
	}
	return false
}
