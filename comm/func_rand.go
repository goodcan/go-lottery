package comm

import (
	"math/rand"
	"time"
)

// 得到一个随机数
func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}
