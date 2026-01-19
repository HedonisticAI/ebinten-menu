package engine

import (
	"ebinten-menus/errors_types"
	"ebinten-menus/menu"
)

type MenuEngine interface {
	Check() //run every update
	Draw()  //run every draw
	Add(menu.MenuUC) errors_types.ID
	Show(errors_types.ID)
}
