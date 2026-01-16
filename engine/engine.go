package engine

import (
	consts "ebinten-menus/const"
	"ebinten-menus/menu"
)

type MenuEngine interface {
	Check()
	Draw()
	Add(menu.MenuUC)
	Show(consts.ID)
}
