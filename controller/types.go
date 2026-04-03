package controller

type Type int
type Status bool

// unified constoller response, Data wil usually be according to type
type ControllerResponse struct {
	Status Status
	Type   Type
	Data   interface{}
}
