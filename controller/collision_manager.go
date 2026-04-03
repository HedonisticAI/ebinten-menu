package controller

import "ebinten-menus/types"

type Controller interface {
	Response() (*ControllerResponse, error)
	SetStatus(Status)
}

type ControllerEngine interface {
	CreateNew(Data interface{}, Type int) Controller
}

type Data interface {
	int | types.ID
}
