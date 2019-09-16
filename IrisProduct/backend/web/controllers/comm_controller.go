package controllers

import "go/types"

type IAppController interface {
	NewService() interface{}
	GetApiPath() string
	GetControllerType() *types.Type
}
