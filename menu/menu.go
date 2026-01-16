package menu

import (
	"ebinten-menus/button"
	"ebinten-menus/format"
	"ebinten-menus/positioner"
)

type Menu struct {
	Buttons    []button.Button
	Format     format.Format
	Positioner positioner.Positioner
}
