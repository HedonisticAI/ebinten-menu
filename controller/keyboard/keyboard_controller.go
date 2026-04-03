package keyboard

import "ebinten-menus/controller"

type KeyboardController struct {
	Status   controller.Status
	Keyboard keyboard
}

func (K *KeyboardController) SetStatus(Status controller.Status) {
	K.Keyboard.SetActive(Status)
}

func (K *KeyboardController) Response() (*controller.ControllerResponse, error) {
	var Res controller.ControllerResponse
	Res.Type = controller.TypeKeyboard
	if K.Keyboard.IsActive() == controller.StatusUnactive {
		return nil, controller.ErrUnactive
	}
	Res.Status = K.Keyboard.IsActive()
	return &Res, nil
}

//dillemma here - should one controller control one button or one menu
//
// book - desolation(?)
// Смерть все еще существует в рамках разложенных сознаний иных индивидов, не?
