package conf

import (
	"sync"
	"time"
)

const (
	SysTimeForm      = "2006-01-02 15:04:05" // 时间转换格式
	SysTimeFormShort = "2006-01-02"          // 日期转换格式
	IpLimitMax       = 500                   // 相同 IP 参与数
	IpPrizeMax       = 10                    // 相同 IP 中奖数
	UserPrizeMax     = 3000                  // 同一用户宗匠数
)

const (
	GiftTypeVirtual   = iota // 虚拟币
	GiftTypeCodeSame         // 虚拟币，相同的码
	GiftTypeCodeDiff         // 虚拟券，不同的码
	GiftTypeGiftSmall        // 实物小奖
	GiftTypeGiftLarge        // 实物大奖
)

var (
	SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	SignSecret         = []byte("1234567890abcedfg")
	CookieSecret       = "hellolottery"
	LoginUser          = new(sync.Map)
)
