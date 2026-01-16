package button

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button interface {
	TriggerPressed()
	Form(ebiten.Image)
}
