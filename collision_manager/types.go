package collision_manager

const StatusCollided = 1
const StatusInteracted = 2
const StatusIdle = 3
const StatusActive = true
const StatusUnactive = false
const TypeKeyboard = 1
const TypeMouse = 2
const TypeCustom = 3

type ControllerResponse struct {
	Status   int
	isActive bool
	Type     int
	Data     map[string]interface{}
}
