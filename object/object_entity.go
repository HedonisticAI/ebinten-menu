package object

import (
	"ebinten-menus/controller"
	"ebinten-menus/drawer"
)

type Object_Entity interface {
	GetDrawer() drawer.Drawer
	GetCollisionManagers() []controller.Controller
	NewManager(controller.Controller) error
	ReplaceDrawer(drawer.Drawer)
}
