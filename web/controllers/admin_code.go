package controllers

import (
	"github.com/kataras/iris"

	"../../services"
)

type AdminCodeController struct {
	Ctx         iris.Context
	ServiceCode services.CodeService
}
