package button

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button_Entity interface {
	TriggerPressed()
	Form(ebiten.Image)
	Check() bool
}
