package utils

import (
	"fmt"
	"log"
	"math"

	"../comm"
	"../dataSource"
)

const ipFrameSize = 2

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
