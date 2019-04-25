package comm

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"

	"go-lottery/conf"
)

// addSlashes 函数返回在预定义字符之前添加反斜杠的字符串
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func AddSlashes(s string) string {
	tmpRune := []rune{}
	sRune := []rune(s)

	for _, ch := range sRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
		default:
			tmpRune = append(tmpRune, ch)
		}
	}

	return string(tmpRune)
}

// stripSlashes 函数删除由 addSlashes 函数添加的反斜杠
func StripSlashes(s string) string {
	dstRune := []rune{}
	sRune := []rune(s)
	sLength := len(sRune)

	for i := 0; i < sLength; i++ {
		if sRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, sRune[i])
	}

	return string(dstRune)
}

// 将字符串的 IP 转化为数字
func Ip4ToInt(ip string) int64 {
	bits := strings.Split(ip, ".")

	if len(bits) == 4 {
		b0, _ := strconv.Atoi(bits[0])
		b1, _ := strconv.Atoi(bits[1])
		b2, _ := strconv.Atoi(bits[2])
		b3, _ := strconv.Atoi(bits[3])

		var sum int64

		sum += int64(b0) << 24
		sum += int64(b1) << 16
		sum += int64(b2) << 8
		sum += int64(b3)

		return sum
	} else {
		return 0
	}
}

// 得到当前时间到下一天零点的延时
func NextDayDuration() time.Duration {
	now := NowTime()
	year, month, day := now.Add(time.Hour * 24).Date()
	next := time.Date(year, month, day, 0, 0, 0, 0, conf.SysTimeLocation)
	return next.Sub(now)
}

// 从接口类型安全获取到 int64
func GetInt64(i interface{}, d int64) int64 {
	if i == nil {
		return d
	}

	switch i.(type) {
	case string:
		num, err := strconv.Atoi(i.(string))

		if err != nil {
			return d
		} else {
			return int64(num)
		}
	case []byte:
		bits := i.([]byte)
		if len(bits) == 0 {
			return int64(binary.LittleEndian.Uint64(bits))
		} else {
			num, err := strconv.Atoi(string(bits))
			if err != nil {
				return d
			} else {
				return int64(num)
			}
		}
	case uint:
		return int64(i.(uint))
	case uint8:
		return int64(i.(uint8))
	case uint16:
		return int64(i.(uint16))
	case uint32:
		return int64(i.(uint32))
	case uint64:
		return int64(i.(uint64))
	case int:
		return int64(i.(int))
	case int8:
		return int64(i.(int8))
	case int16:
		return int64(i.(int16))
	case int32:
		return int64(i.(int32))
	case int64:
		return i.(int64)
	case float32:
		return int64(i.(float32))
	case float64:
		return int64(i.(float64))
	}

	return d
}

// 从接口类型安全获取到字符串类型
func GetString(s interface{}, d string) string {
	if s == nil {
		return d
	}

	switch s.(type) {
	case string:
		return s.(string)
	case []byte:
		return string(s.([]byte))
	}

	return fmt.Sprintf("%s", s)
}

// 从 map 中获得指定的 key
func GetInt64FromMap(m map[string]interface{}, key string, d int64) int64 {
	data, ok := m[key]

	if !ok {
		return d
	}

	return GetInt64(data, d)
}

// 从 map 中获得指定的 key
func GetInt64FromStringMap(m map[string]string, key string, d int64) int64 {
	data, ok := m[key]

	if !ok {
		return d
	}

	return GetInt64(data, d)
}

// 从 map 中获得指定的 key
func GetStringFromMap(m map[string]interface{}, key string, d string) string {
	data, ok := m[key]

	if !ok {
		return d
	}

	return GetString(data, d)
}

// 从 map 中获得指定的 key
func GetStringFromStringMap(m map[string]string, key string, d string) string {
	data, ok := m[key]

	if !ok {
		return d
	}

	return GetString(data, d)
}

// 从 上下文获取，返回结果
func FromCtxGetResult(ctx iris.Context) *conf.Result {
	rs := ctx.Values().Get("result")
	switch rs.(type) {
	case *conf.Result:
		return rs.(*conf.Result)
	default:
		return nil
	}
}
