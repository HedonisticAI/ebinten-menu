package button

import (
	"ebinten-menus/collision_manager"
)

type Button struct {
	Type               int
	Collision_managers []collision_manager.Collision_manager
	//every button can be activated thought keyboard button, mouse click or gamepad(?)
	//idk how to implement gamepad
}
