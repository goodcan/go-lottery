package utils

import (
	"fmt"
	"log"
	"math"
	"time"

	"go-lottery/comm"
	"go-lottery/dataSource"
)

const ipFrameSize = 2

func init() {
	resetGroupIpList()
}

func resetGroupIpList() {
	log.Println("ip_day_lucky.resetGroupIpList start")
	redisDB := dataSource.RedisInstCache()
	for i := 0; i < ipFrameSize; i++ {
		key := fmt.Sprintf("day_ips_%d", i)
		_, _ = redisDB.Do("DEL", key)
	}
	log.Println("ip_day_lucky.resetGroupIpList stop")

	// IP 当天的统计数，零点的时候归零，设置定时器
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupIpList)
}

func IncrIpLuckyNum(strIp string) int64 {
	ip := comm.Ip4ToInt(strIp)
	i := ip % ipFrameSize
	key := fmt.Sprintf("day_ips_%d", i)
	redisDB := dataSource.RedisInstCache()
	rs, err := redisDB.Do("HINCRBY", key, ip, 1)
	if err != nil {
		log.Println("ip_day_lucky redis HINCRBY error ", err)
		return math.MaxInt64
	}
	return rs.(int64)
}
