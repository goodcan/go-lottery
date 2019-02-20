package server

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strings"
)

// 抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

// 首页 奖品信息
// GET http://localhost:8080/
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br/>%s",
		strings.Join(prizeList, "<br/>"))
}

// 奖品概率
// GET http://localhost:8080/debug
func (c *lotteryController) GetDebug() string {
	c.Ctx.Header("Content-type", "text/html")
	return fmt.Sprintf("抽奖概率：<br/>%s",
		strings.Join(func(rateList []Prate) []string {
			rs := []string{}
			for _, v := range rateList {
				rs = append(rs, fmt.Sprintf("%+v", v))
			}
			return rs
		}(rateList), "<br/>"))
}

func NewApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func Run() {
	app := NewApp()
	app.Run(iris.Addr(":8080"))
}
