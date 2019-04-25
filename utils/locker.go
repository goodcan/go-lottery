package utils

import (
	"fmt"

	"go-lottery/dataSource"
)

func getLuckyLockKye(uid int) string {
	return fmt.Sprintf("lucky_lock_%s", uid)
}

func LockLucky(uid int) bool {
	key := getLuckyLockKye(uid)
	redisDB := dataSource.RedisInstCache()
	rs, _ := redisDB.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

func UnLockLucky(uid int) bool {
	key := getLuckyLockKye(uid)
	redisDB := dataSource.RedisInstCache()
	rs, _ := redisDB.Do("DEL", key)
	if rs == "OK" {
		return true
	} else {
		return false
	}
}
