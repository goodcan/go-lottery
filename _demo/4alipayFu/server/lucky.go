package server

import (
	"math/rand"
	"time"
)

func luckyCode() int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(rateMax)
}

// 记录用户获奖信息
func saveLuckyData(code int32, sendData string, g *gift) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s \n",
		code, g.id, g.name, g.link, sendData)
}
