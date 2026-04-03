package controller

import "errors"

// errors
var ErrNotSupported = errors.New("controller type not supported")
var ErrUnactive = errors.New("this controller is unactive")

// statuses
const StatusActive = true
const StatusUnactive = false

// controller types
const TypeKeyboard = 1
const TypeMouse = 0
const TypeGamepad = 2
const TypeCustom = 3
