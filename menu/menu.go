package menu

import (
	"ebinten-menus/format"
	"ebinten-menus/object"
)

type Menu struct {
	Buttons []object.Object_Entity
	Format  format.Format
}
