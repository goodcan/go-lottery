/**
 * 1 即开即得星 http://localhost:8080/
 * 2 双色球自选型 http://localhost:8080/prize
 */
package server

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotteryController struct {
	Ctx iris.Context
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

// 即开即得型型 http://localhost:8080/
func (c *lotteryController) Get() string {
	var prize string

	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)

	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >= 4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf(
			"尾号为1蝴蝶一等奖<br/>"+
				"尾号为2/3获得二等奖<br/>"+
				"尾号为4/5/6获得三等奖"+
				"cdoe+%d<br/>"+
				"很遗憾，没有获奖", code)
	}

	return fmt.Sprintf(
		"尾号为1蝴蝶一等奖<br/>"+
			"尾号为2/3获得二等奖<br/>"+
			"尾号为4/5/6获得三等奖"+
			"cdoe+%d<br/>"+
			"恭喜您获得%s", code, prize)
}

// 双色球自选型  http://localhost:8080/prize
func (c *lotteryController) GetPrize() string {

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	var prize [7]int

	// 6 个红色球，1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}

	// 最后一位的蓝色球，1-16
	prize[6] = r.Intn(16) + 1

	return fmt.Sprintf("今日开奖号码是：%v", prize)
}
