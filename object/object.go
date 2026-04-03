package object

import (
	"ebinten-menus/controller"
	"ebinten-menus/drawer"
)

type Object struct {
	Controllers []controller.Controller
	//every button can be activated thought keyboard button, mouse click or gamepad(?)
	//idk how to implement gamepad
	//and i don't have to think about it

	Drawer drawer.Drawer
}

func (O *Object) GetDrawer() drawer.Drawer {
	return O.Drawer
}

func (O *Object) GetCollisionManagers() []controller.Controller {
	return O.GetCollisionManagers()
}

func (O *Object) ReplaceDrawer(D drawer.Drawer) error {
	O.Drawer = D
	return nil
}
