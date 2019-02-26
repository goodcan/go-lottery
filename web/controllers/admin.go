package controllers

import (
	"github.com/kataras/iris"

	"../../comm"
	"../../services"
)

type AdminController struct {
	Ctx              iris.Context
	ServiceUserDay   services.UserDayService
	ServiceCode      services.CodeService
	ServiceGift      services.GiftService
	ServiceResult    services.ResultService
	ServiceBlackIp   services.BlackIpService
	ServiceBlackUser services.BlackUserService
}

func (this *AdminController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)

	//uri := "http://" + this.Ctx.Host() + "/admin"

	rs.Data = map[string]interface{}{
		"title": "管理后台",
	}

	this.Ctx.Next()
}
