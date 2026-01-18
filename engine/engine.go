package engine

import (
	"ebinten-menus/consts_types"
	"ebinten-menus/menu"
)

type MenuEngine interface {
	Check() //run every update
	Draw()  //run every draw
	Add(menu.MenuUC) consts_types.ID
	Show(consts_types.ID)
}
