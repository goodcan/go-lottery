/**
 * 设置红包
 * curl "http://localhost:8080/set?uid=1&money=100&num=100"
 * 抢红包
 * curl "http://localhost:8080/get?uid=1&id2994023272"
 * 并发压力测试
 * wrk -t 10 -c 10 -d 5 http://localhost:8080/set?uid=1&money=100&num=100
 * packageList 中 map 和 map 中的 list 都要处理线程安全
 */
package server

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// 红包列表
// map 线程不安全, *sync.Map 线程安全
// var packageList = make(map[uint32][]uint)
// 使用 *sync.Map 对应的使用方法也会改变
var packageList = new(sync.Map)

type lotteryController struct {
	Ctx iris.Context
}

// 返回全部红包地址
// GET http://localhost:8080/
func (c *lotteryController) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)

	//for id, list := range packageList {
	//	var money int
	//	for _, v := range list {
	//		money += int(v)
	//	}
	//	rs[id] = [2]int{len(list), money}
	//}

	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]uint)

		var money int
		for _, v := range list {
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}

		return true
	})

	return rs
}

// 发红包
// GET http://localhost:8080/set?uid=1&money=100&num=100
func (c *lotteryController) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")

	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("参数格式异常，errUid=%d, errMoney=%f, errNum=%d",
			errUid, errMoney, errNum)
	}

	moneyTotal := int(money * 100)

	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%d, num=%d",
			uid, moneyTotal, num)
	}

	// 金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 随机分配的最大值，例如 10 元，最大的红包为 5.5 元
	rMax := 0.55
	if num >= 1000 {
		rMax = 0.01
	} else if num >= 100 {
		rMax = 0.1
	} else if num >= 10 {
		rMax = 0.3
	}

	// 红包数量
	list := make([]uint, num)

	// 剩余红包数
	leftNum := num

	// 剩余金额
	leftMoney := moneyTotal

	// 大循环开始，分配金额给到每一个红包
	for leftNum > 0 {

		// 最后一个红包，剩余的全部金额都给它
		if leftNum == 1 {
			list[num-1] = uint(leftMoney)
			break
		}

		// 剩余金额等于剩余个数，则每人只能分到一分钱
		if leftMoney == leftNum {
			for i := num - leftNum; i < leftNum; i++ {
				list[i] = 1
			}
			break
		}

		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}

		list[num-leftNum] = uint(m)
		leftNum--
		leftMoney -= m

	}

	// 红包唯一 ID
	id := r.Uint32()
	//packageList[id] = list
	packageList.Store(id, list)

	// 返回抢红包的URL
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}

// 抢红包
// GET http://localhost:8080/get?id=1&uid=1&num=100
func (c *lotteryController) GetGet() string {
	id, errId := c.Ctx.URLParamInt("id")
	uid, errUid := c.Ctx.URLParamInt("uid")

	if errUid != nil || errId != nil {
		return fmt.Sprintf("参数格式错误，uid=%d, id=%d", uid, id)
	}

	if uid < 1 || id < 1 {
		return fmt.Sprintf("参数数据异常，uid=%d, id=%d", uid, id)
	}

	//list, ok := packageList[uint32(id)]
	list1, ok := packageList.Load(uint32(id))
	list := list1.([]uint)

	l := len(list)

	if !ok || l < 1 {
		return fmt.Sprintf("红包不存在")
	}

	// 构造一个抢红包任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback: callback}
	// 发送任务
	chTasks := chTaskList[id%taskNum]
	chTasks <- t
	// 接受返回结果
	money := <-t.callback

	if money <= 0 {
		return fmt.Sprintf("很遗憾，没有抢到红包")
	} else {
		return fmt.Sprintf("恭喜你抢到一个红包，金额为：%.2f", float64(money)*0.01)
	}

	/*
		 * 最原始的线程不安全写法

		// 分配随机数
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := r.Intn(l)
		money := list[i]

		// 更新红包列表中的信息
		if l > 1 {
			if i == l-1 {
				//packageList[uint32(id)] = list[:i]
				packageList.Store(uint32(id), list[:i])
			} else if i == 0 {
				//packageList[uint32(id)] = list[1:]
				packageList.Store(uint32(id), list[1:])
			} else {
				//packageList[uint32(id)] = append(list[:i], list[i+1:]...)
				packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
			}
		} else {
			//delete(packageList, uint32(id))
			packageList.Delete(uint32(id))
		}

	*/
}

func NewApp() *iris.Application {
	app := iris.Default()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	for i := 0; i < taskNum; i++ {
		chTaskList[i] = make(chan task)
		go fetchPackageListMoney(chTaskList[i])
	}

	return app
}

func Run() {
	app := NewApp()
	app.Run(iris.Addr(":8080"))
}
