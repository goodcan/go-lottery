package controllers

import (
	"../../comm"
	"../../utils"
)

// GET http://localhost:8080/lucky
func (this *IndexController) GetLucky() {
	rs := comm.FromCtxGetResult(this.Ctx)

	loginUser := comm.GetLoginUser(this.Ctx.Request())

	// 1 验证登录用户
	if loginUser == nil || loginUser.Uid < 1 {
		rs.SetError(101, "请先登录，再来抽奖")
		this.Ctx.Next()
	}

	// 2 用户抽奖分布式锁定
	ok := utils.LockLucky(loginUser.Uid)
	if !ok {
		rs.SetError(102, "正在抽奖，请稍后重试")
		this.Ctx.Next()
	} else {
		defer utils.UnLockLucky(loginUser.Uid)
	}

	// 3 验证用户今日参与次数

	// 4 验证 IP 今日的参与次数

	// 5 验证 IP 黑名单

	// 6 验证用户黑名单

	// 7 获得抽奖编码

	// 8 匹配奖品是否中奖

	// 9 有限制奖品发放

	// 10 不用编码的优惠券的发放

	// 11 记录中奖记录

	// 12 返回抽奖结果
	this.Ctx.Next()
}
