package controllers

import (
	"go-lottery/comm"
	"go-lottery/models"
)

func (this *IndexController) prizeLarge(ip string,
	loginUser *models.LoginUser,
	userInfo *models.BlackUser,
	blackIpInfo *models.BlackIp) {
	now := comm.NowUnix()
	blackTime := 30 * 86400

	// 更新用户黑名单信息
	if userInfo == nil || userInfo.Id <= 0 {
		userInfo := &models.BlackUser{
			Id:         loginUser.Uid,
			Uid:        loginUser.Uid,
			Username:   loginUser.Username,
			BlackTime:  comm.StampToTime(now + blackTime),
			SysCreated: comm.StampToTime(now),
			SysIP:      ip,
		}
		this.ServiceBlackUser.Insert(userInfo)
	} else {
		userInfo.Id = loginUser.Uid
		userInfo.BlackTime = comm.StampToTime(now + blackTime)
		userInfo.SysUpdated = comm.StampToTime(now)
		this.ServiceBlackUser.Update(userInfo, nil)
	}

	// 更新 IP 黑名单
	if blackIpInfo == nil || blackIpInfo.Id <= 0 {
		blackIpInfo = &models.BlackIp{
			Ip:         ip,
			BlackTime:  comm.StampToTime(now + blackTime),
			SysCreated: comm.StampToTime(now),
		}
		this.ServiceBlackIp.Insert(blackIpInfo)
	} else {
		blackIpInfo.BlackTime = comm.StampToTime(now + blackTime)
		blackIpInfo.SysUpdated = comm.StampToTime(now)
		this.ServiceBlackIp.Update(blackIpInfo, nil)
	}
}
