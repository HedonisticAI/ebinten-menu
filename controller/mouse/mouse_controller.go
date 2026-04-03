package mouse

import (
	"ebinten-menus/controller"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MouseController struct {
	Status        controller.Status
	Form          MouseForm
	ActiveButtons []ebiten.MouseButton
}

// If Button is in list, its status will be checked every tick
func (M *MouseController) Response() (*controller.ControllerResponse, error) {
	var Res controller.ControllerResponse
	var Data *MouseData
	Res.Status = M.Status
	if M.Status == controller.StatusUnactive {
		return nil, controller.ErrUnactive
	}
	for _, Button := range M.ActiveButtons {
		Data.Buttons[Button] = NoAction
		if inpututil.IsMouseButtonJustPressed(Button) {
			Data.Buttons[Button] = Pressed
		}
		if inpututil.IsMouseButtonJustReleased(Button) {
			Data.Buttons[Button] = Released
		}
	}
	Data.Collision = M.Form.Collison()
	Res.Data = Data
	return &Res, nil
}
func (M *MouseController) SetStatus(Status controller.Status) {
	M.Status = Status
}
