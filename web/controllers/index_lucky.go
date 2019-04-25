package controllers

import (
	"fmt"
	"log"

	"go-lottery/comm"
	"go-lottery/conf"
	"go-lottery/models"
	"go-lottery/utils"
)

// GET http://localhost:8080/lucky
func (this *IndexController) GetLucky() {
	rs := comm.FromCtxGetResult(this.Ctx)

	loginUser := comm.GetLoginUser(this.Ctx.Request())

	// 1 验证登录用户
	if loginUser == nil || loginUser.Uid < 1 {
		rs.SetError(101, "请先登录，再来抽奖")
		this.Ctx.Next()
		return
	}

	// 2 用户抽奖分布式锁定
	ok := utils.LockLucky(loginUser.Uid)
	if !ok {
		rs.SetError(102, "正在抽奖，请稍后重试")
		this.Ctx.Next()
		return
	} else {
		defer utils.UnLockLucky(loginUser.Uid)
	}

	// 3 验证用户今日参与次数
	ok = this.checkUserDay(loginUser.Uid)
	if !ok {
		rs.SetError(103, "今日的抽奖次数已用完，明天再来吧")
		this.Ctx.Next()
		return
	}

	// 4 验证 IP 今日的参与次数
	ip := comm.ClientIp(this.Ctx.Request())
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs.SetError(104, "相同 IP 参与次数太多，明天再来吧")
		this.Ctx.Next()
		return
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

	not_prize_msg := "很遗憾，没有中奖，请下次再试"

	// 8 匹配奖品是否中奖
	prizeGift := this.prize(prizeCode, limitBlack)

	if prizeGift == nil || prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		rs.SetError(205, not_prize_msg)
		this.Ctx.Next()
		return
	}

	// 9 有限制奖品发放
	if prizeGift.PrizeNum > 0 {
		ok = utils.PrizeGift(prizeGift.Id, this.ServiceGift)
		if !ok {
			rs.SetError(207, not_prize_msg)
			this.Ctx.Next()
			return
		}
	}

	// 10 不用编码的优惠券的发放
	if prizeGift.Gtype == conf.GiftTypeCodeDiff {
		code := utils.PrizeCodeDiff(prizeGift.Id, this.ServiceCode)
		if code == "" {
			rs.SetError(208, not_prize_msg)
			this.Ctx.Next()
			return
		}
		prizeGift.Gdata = code
	}

	// 11 记录中奖记录
	result := models.Result{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        loginUser.Uid,
		Username:   loginUser.Username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowTime(),
		SysStatus:  0,
		SysIP:      ip,
	}

	err := this.ServiceResult.Insert(&result)

	if err != nil {
		log.Println("index_lucky.GetLucky ServiceResult.Insert",
			result, "error=", err)
		rs.SetError(209, not_prize_msg)
		this.Ctx.Next()
		return
	}

	// 如果是实物大奖需要将用户 IP 设置成黑名单一段时间
	if prizeGift.Gtype == conf.GiftTypeGiftLarge {
		this.prizeLarge(ip, loginUser, userInfo, blackIpInfo)
	}

	// 12 返回抽奖结果
	rs.Data = prizeGift
	this.Ctx.Next()
}
