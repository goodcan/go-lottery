package controllers

import (
	"go-lottery/comm"
	"go-lottery/models"
)

func (this *IndexController) checkBlackUser(uid int) (*models.BlackUser, bool) {
	info := this.ServiceBlackUser.GetByUid(uid)
	if info != nil && comm.TimeToStamp(info.BlackTime) > comm.NowUnix() {
		// 黑名单存在，并且有效
		return info, false
	} else {
		return nil, true
	}
}
