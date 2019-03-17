package server

import (
	"math/rand"
	"time"
)

func luckyCode() int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(rateMax)
}

// 虚拟币
func sendCoin(g *gift) (bool, string) {
	if g.total == 0 {
		// 数量无线
		return true, g.data
	} else if g.left > 0 {
		// 还有剩余
		g.left -= 1
		return true, g.data
	} else {
		return true, "奖品已发完"
	}
}

// 不同值的优惠券
func sendCoupon(g *gift) (bool, string) {
	if g.left > 0 {
		// 还有剩余
		g.left -= 1

		if g.left > len(g.dataList) {
			return true, "奖品已发完"
		} else {
			return true, g.dataList[g.left]
		}

	} else {
		return true, "奖品已发完"
	}
}

var sendCouponFix = sendCoin
var sendRealSmall = sendCoin
var sendRealLarge = sendCoin

// 记录用户获奖信息
func saveLuckyData(code int32, sendData string, g *gift) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s, left=%d \n",
		code, g.id, g.name, g.link, sendData, g.left)
}
