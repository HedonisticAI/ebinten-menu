package engine

import (
	"ebinten-menus/menu"
	types "ebinten-menus/types"
)

type MenuEngine interface {
	Check() //run every update
	Draw()  //run every draw
	Add(menu.MenuUC) types.ID
	Show(types.ID)
}
