package comm

import (
	"time"

	"go-lottery/conf"
)

// 当前时间的时间戳
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}

func NowTime() time.Time {
	return time.Now().In(conf.SysTimeLocation)
}

// 将 unix 时间戳格式化为 yyyymmdd HH:II:SS 格式字符串
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeForm)
	} else {
		return NowTime().Format(conf.SysTimeForm)
	}
}

// 将 unix 时间戳格式化为 yyyymmdd 格式字符串
func FormatFromUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeFormShort)
	} else {
		return NowTime().Format(conf.SysTimeFormShort)
	}
}

func StrShortToTime(s string) time.Time {
	t, _ := time.Parse(conf.SysTimeFormShort, s)
	return t.In(conf.SysTimeLocation)
}

// 将字符串转成时间
func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation(conf.SysTimeForm, s, conf.SysTimeLocation)
}

func StampToTime(u int) time.Time {
	return time.Unix(int64(u), 0).In(conf.SysTimeLocation)
}

func TimeToStamp(t time.Time) int {
	return int(t.In(conf.SysTimeLocation).Unix())
}
