package object

import (
	"ebinten-menus/collision_manager"
	"ebinten-menus/drawer"
)

type Object_Entity interface {
	GetDrawer() drawer.Drawer
	GetCollisionManagers() []collision_manager.Controller
	NewManager(collision_manager.Controller) error
	ReplaceDrawer(drawer.Drawer)
}
