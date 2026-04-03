package mouse

import "github.com/hajimehoshi/ebiten/v2"

const Pressed = 0
const Released = 1
const NoAction = -1

type MouseData struct {
	Buttons   map[ebiten.MouseButton]int
	Collision bool
}
