package menu

import (
	"ebinten-menus/button"
	"ebinten-menus/format"
)

type Menu struct {
	Buttons []button.Button_Entity
	Format  format.Format
}
