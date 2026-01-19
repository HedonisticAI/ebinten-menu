package button

import (
	"ebinten-menus/collision_manager"
	"ebinten-menus/drawer"
)

type Button struct {
	Collision_managers []collision_manager.Collision_manager
	//every button can be activated thought keyboard button, mouse click or gamepad(?)
	//idk how to implement gamepad
	Drawer drawer.Drawer
}
