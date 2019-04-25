package controllers

import (
	"github.com/kataras/iris"

	"go-lottery/comm"
)

type AdminController struct {
	Ctx iris.Context
}

func (this *AdminController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)

	uri := "http://" + this.Ctx.Host() + "/admin"
	uri_result := uri + "/result"
	uri_black_user := uri + "/blackUser"
	uri_black_ip := uri + "/blackIp"

	rs.Data = map[string]interface{}{
		"title":                  "管理后台",
		"admin_result":           uri_result,
		"admin_result_delete":    uri_result + "/delete",
		"admin_black_user":       uri_black_user,
		"admin_black_user_black": uri_black_user + "/black",
		"admin_black_ip":         uri_black_ip,
		"admin_black_ip_black":   uri_black_ip + "/black",
	}

	this.Ctx.Next()
}
