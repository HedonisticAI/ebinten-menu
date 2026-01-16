package consts

import "errors"

const TypeKeySControl = 1
const TypeGamepadControl = 2
const TypeMouseControl = 3

var ErrOutOfBounds = errors.New("ouf of bounds")

type UserClick struct {
}
type ID int
