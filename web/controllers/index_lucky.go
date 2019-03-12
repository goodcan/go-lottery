package controllers

import (
	"../../comm"
	"../../conf"
	"../../models"
	"../../utils"
	"fmt"
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
	ok = this.checkUserDay(loginUser.Uid)
	if !ok {
		rs.SetError(103, "今日的抽奖次数已用完，明天再来吧")
		this.Ctx.Next()
	}

	// 4 验证 IP 今日的参与次数
	ip := comm.ClientIp(this.Ctx.Request())
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs.SetError(104, "相同 IP 参与次数太多，明天再来吧")
		this.Ctx.Next()
	}

	// 黑名单
	limitBlack := false
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}

	// 5 验证 IP 黑名单
	var blackIpInfo *models.BlackIp
	if !limitBlack {
		blackIpInfo, ok = this.checkBlackIp(ip)
		if !ok {
			fmt.Println("黑名单中的 IP", ip, blackIpInfo.BlackTime)
			limitBlack = true
		}
	}

	// 6 验证用户黑名单
	var userInfo *models.BlackUser
	if !limitBlack {
		userInfo, ok = this.checkBlackUser(loginUser.Uid)
		if !ok {
			fmt.Println("黑名单中的用户", loginUser.Uid, userInfo.BlackTime)
			limitBlack = true
		}
	}

	// 7 获得抽奖编码
	prizeCode := comm.RandInt(10000)

	// 8 匹配奖品是否中奖
	prizeGift := this.prize(prizeCode, limitBlack)

	if prizeGift == nil || prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		rs.SetError(205, "很遗憾，没有中奖，请下次再试")
		this.Ctx.Next()
	}

	// 9 有限制奖品发放

	// 10 不用编码的优惠券的发放

	// 11 记录中奖记录

	// 12 返回抽奖结果
	this.Ctx.Next()
}
