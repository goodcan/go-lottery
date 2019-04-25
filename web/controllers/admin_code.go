package controllers

import (
	"github.com/kataras/iris"

	"go-lottery/services"
)

type AdminCodeController struct {
	Ctx         iris.Context
	ServiceCode services.CodeService
}
