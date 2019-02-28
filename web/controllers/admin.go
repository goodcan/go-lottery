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

	uri := "http://" + this.Ctx.Host() + "/admin"
	uri_result := uri + "/result"
	uri_black_user := uri + "/blackUser"

	rs.Data = map[string]interface{}{
		"title":                  "管理后台",
		"admin_result":           uri_result,
		"admin_result_delete":    uri_result + "/delete",
		"admin_black_user":       uri_black_user,
		"admin_black_user_black": uri_black_user + "/black",
	}

	this.Ctx.Next()
}
