package server

import (
	"log"
	"os"
	"sync"
)

// 奖品类型，枚举值 iota 从 0 开始
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同券
	giftTypeCouponFix        // 相同的券
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int      // 奖品ID
	name     string   // 奖品名称
	pic      string   // 奖品连接
	link     string   // 奖品连接
	gtype    int      // 奖品类型
	data     string   // 奖品数据（特定的配置信息）
	dataList []string // 奖品数据集合（不同的优惠券的编码
	total    int      // 总数，0 为无限
	left     int      // 剩余数量
	isUse    bool     // 是否可用
	rate     int      // 中奖概率，万分之 N, 0-9999
	rateMin  int      // 大于等于最小中奖编码
	rateMax  int      // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10000

var logger *log.Logger
var mu sync.Mutex

var giftList [5]*gift

// 初始化日志
func initLog() {
	f, _ := os.Create("./_demo/3wechatShake/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

// 初始化奖品列表
func initGift() {
	giftList[0] = &gift{
		id:       1,
		name:     "手机大奖",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealLarge,
		data:     "",
		dataList: nil,
		total:    10,
		left:     10,
		isUse:    true,
		rate:     1,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[1] = &gift{
		id:       2,
		name:     "充电器",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealSmall,
		data:     "",
		dataList: nil,
		total:    5,
		left:     5,
		isUse:    false,
		rate:     10,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[2] = &gift{
		id:       3,
		name:     "优惠券满200减50元",
		pic:      "",
		link:     "",
		gtype:    giftTypeCouponFix,
		data:     "mall-coupon-2019",
		dataList: nil,
		total:    50,
		left:     50,
		isUse:    false,
		rate:     500,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[3] = &gift{
		id:       4,
		name:     "直降优惠券50元",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		dataList: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    10,
		left:     10,
		isUse:    false,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[4] = &gift{
		id:       5,
		name:     "虚拟币",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoin,
		data:     "10金币",
		dataList: nil,
		total:    5,
		left:     5,
		isUse:    false,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}

	// 数据整理，中奖区间数据
	rateStart := 0
	for _, data := range giftList {
		if !data.isUse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate

		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}
