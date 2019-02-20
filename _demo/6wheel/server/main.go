/*
 # 压力测试：
 * wrk -t 10 -c 100 -d 5 "http://localhost:8080/prize"
*/
package server

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
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
			var rs []string
			for _, v := range rateList {
				rs = append(rs, fmt.Sprintf("%+v", v))
			}
			return rs
		}(rateList), "<br/>"))
}

func (c *lotteryController) GetPrize() string {
	// 第一步，抽奖，根据随机数匹配奖品
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(10000)

	var myPrize string
	var prizeRate *Prate

	// 从奖品列表匹配是否中奖
	for i, prize := range prizeList {
		rate := &rateList[i]

		// 满足中奖条件
		if code >= rate.CodeA && code <= rate.CodeB {
			myPrize = prize
			prizeRate = rate
			break
		}
	}

	if myPrize == "" {
		return "很遗憾，再来一次吧"
	}

	// 第二步，中奖后开始发奖
	// 无限量奖品
	if prizeRate.Total == 0 {
		logger.Println("奖品", myPrize)
		return myPrize
	} else if *prizeRate.Left > 0 {
		left := atomic.AddInt32(prizeRate.Left, -1)

		if left >= 0 {
			logger.Println("奖品", myPrize)
			return myPrize
		} else {
			return "很遗憾，再来一次吧"
		}
	} else {
		return "很遗憾，再来一次吧"
	}
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
