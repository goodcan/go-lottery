package server

// 奖品中奖概率
type Prate struct {
	Rate  int // 万分之 N 的中奖概率
	Total int // 总数量限制，0 表示无限数量
	CodeA int // 中奖概率起始编码（包含）
	CodeB int // 中奖概率终止编码（包含）
	Left  int // 剩余数
}

// 奖品列表
var prizeList = []string{
	"一等奖，火星单程船票",
	"二等奖，凉嗖嗖南极之旅",
	"三等奖，iPhone一部",
}

// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
var rateList = []Prate{
	{1, 1, 0, 0, 1},
	{2, 2, 1, 2, 1},
	{5, 10, 3, 5, 10},
	{100, 0, 0, 9999, 0},
}
