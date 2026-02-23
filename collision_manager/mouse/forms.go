package mouse

import (
	"ebinten-menus/errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonEllipse struct {
	Active bool
	X      int
	Y      int
	XAxis  int
	YAxis  int
}

func (E *ButtonEllipse) SetActive(Active bool) {
	E.Active = Active
}

func (E *ButtonEllipse) IsActive() bool {
	return E.Active
}

func (E *ButtonEllipse) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, errors.ErrMouseNotSupported
	}
	if E.ellipse_formula(x, y) {
		return true, nil
	}
	return false, nil
}

func (E *ButtonEllipse) ellipse_formula(x int, y int) bool {
	formula := ((x-E.X)^2)/(E.XAxis^2) + ((y-E.Y)^2)/(E.YAxis^2)
	return formula >= 1
}

type ButtonRound struct {
	Active bool
	X      int
	Y      int
	Raduis int
}

func (E *ButtonRound) SetActive(Active bool) {
	E.Active = Active
}

func (E *ButtonRound) IsActive() bool {
	return E.Active
}

func (R *ButtonRound) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, errors.ErrMouseNotSupported
	}
	if (x <= R.X+R.Raduis && x >= R.X-R.Raduis) && (y <= R.Y+R.Raduis && y >= R.Y-R.Raduis) {
		return true, nil
	}
	return false, nil
}

type ButtonRectangle struct {
	Active bool
	X      int // upper left
	Y      int //
	Lenght int
	Width  int
}

func (E *ButtonRectangle) SetActive(Active bool) {
	E.Active = Active
}

func (E *ButtonRectangle) IsActive() bool {
	return E.Active
}
func (B *ButtonRectangle) IsCollision() (bool, error) {
	x, y := ebiten.CursorPosition()
	if x == 0 && y == 0 && ebiten.Tick() != 0 {
		return false, errors.ErrMouseNotSupported
	}
	if (x <= B.X) && (x >= B.X+B.Lenght) && (y <= B.Y) && (y >= B.Y-B.Width) {
		return true, nil
	}
	return false, nil
}
