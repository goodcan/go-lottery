/*
 * 压力测试
 *  wrk -t10 -c10 -d5 http://localhost:8080/lucky
 */
package server

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotteryController struct {
	Ctx iris.Context
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRage(rate)

	result := ""
	for _, data := range giftList {
		result += fmt.Sprintf("%v\n", data)
	}

	return result
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	giftList := giftRage(rate)

	result := map[string]interface{}{}

	result["success"] = false

	for _, data := range giftList {
		if !data.isUse {
			continue
		}

		// 中奖了，抽奖编码在奖品编码范围内
		if data.rateMin <= int(code) && int(code) < data.rateMax {
			sendData := data.pic
			saveLuckyData(code, sendData, &data)
			result["success"] = true
			result["uid"] = uid
			result["id"] = data.id
			result["name"] = data.name
			result["link"] = data.link
			result["data"] = sendData
			break
		}
	}

	if v, ok := result["success"]; ok && v == false {
		result["data"] = "没有中奖"
	}

	return result
}

func NewApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()

	return app
}

func Run() {
	app := NewApp()
	app.Run(iris.Addr(":8080"))

}
