package server

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 奖品连接
	link    string // 奖品连接
	isUse   bool   // 是否可用
	rate    int    // 中奖概率，万分之 N, 0-9999
	rateMin int    // 大于等于最小中奖编码
	rateMax int    // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10

var logger *log.Logger
var mu sync.Mutex

// 初始化日志
func initLog() {
	f, _ := os.Create("./_demo/4alipayFu/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newGift() *[5]gift {
	giftList := new([5]gift)

	giftList[0] = gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[1] = gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[2] = gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[3] = gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[4] = gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		isUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}

	return giftList
}

func giftRage(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)

	rateStart := 0
	for i := range giftList {
		if !giftList[i].isUse {
			continue
		}

		grate := 0

		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}

		giftList[i].rate = grate

		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate

		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}

	//fmt.Printf("giftLsit=%v\n", giftList)

	return giftList
}
