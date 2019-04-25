package controllers

import (
	"go-lottery/comm"
	"go-lottery/models"
)

func (this *IndexController) checkBlackIp(ip string) (*models.BlackIp, bool) {
	info := this.ServiceBlackIp.GetByIp(ip)
	if info != nil && comm.TimeToStamp(info.BlackTime) > comm.NowUnix() {
		// IP 黑名单存在，并且还在黑名单有效期内
		return info, false
	} else {
		return nil, true
	}
}
