package controllers

import (
	"github.com/kataras/iris"

	"../../models"
	"../../services"
)

type IndexController struct {
	Ctx              iris.Context
	ServiceUserDay   services.UserDayService
	ServiceCode      services.CodeService
	ServiceGift      services.GiftService
	ServiceResult    services.ResultService
	ServiceBlackIp   services.BlackIpService
	ServiceBlackUser services.BlackUserService
}

// http://localhost:8080/
func (this *IndexController) Get() string {
	this.Ctx.Header("Content-Type", "text/html")
	return "<h2>welcome to Go lottery</h2><a href='/public/index.html'>开始抽奖</a>"
}

// http://localhost:8080/gifts
func (this *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	dataList := this.ServiceGift.GetAll()
	gifts := make([]models.Gift, 0)

	for _, data := range dataList {
		if data.SysStatus == 0 {
			gifts = append(gifts, data)
		}
	}

	rs["gifts"] = gifts

	return rs
}

// http://localhost:8080/new/prize
func (this *IndexController) GetNewPrize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	// TODO

	return rs
}
