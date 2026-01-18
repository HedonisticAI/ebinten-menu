package mouse

import (
	"ebinten-menus/consts_types"

	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonEllipse struct {
	X     int
	Y     int
	XAxis int
	YAxis int
}

func (E *ButtonEllipse) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, consts_types.ErrMouseNotSupported
	}
	if E.ellipse_formula(x, y) {
		return true, nil
	}
	return false, nil
}

func (E *ButtonEllipse) ellipse_formula(x int, y int) bool {
	formula := ((x-E.X)^2)/(E.XAxis^2) + ((y-E.Y)^2)/E.YAxis ^ 2
	return formula >= 1
}

type ButtonRound struct {
	X      int
	Y      int
	Raduis int
}

func (R *ButtonRound) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, consts_types.ErrMouseNotSupported
	}
	if (x <= R.X+R.Raduis && x >= R.X-R.Raduis) && (y <= R.Y+R.Raduis && y >= R.Y-R.Raduis) {
		return true, nil
	}
	return false, nil
}

type ButtonRectangle struct {
	X      int // upper left
	Y      int //
	Lenght int
	Width  int
}

func (B *ButtonRectangle) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, consts_types.ErrMouseNotSupported
	}
	if (x <= B.X) && (x >= B.X+B.Lenght) && (y <= B.Y) && (y >= B.Y-B.Width) {
		return true, nil
	}
	return false, nil
}
