package controllers

import (
	"go-lottery/conf"
	"go-lottery/models"
)

func (this *IndexController) prize(prizeCode int, limitBlack bool) *models.GiftPrize {
	var prizeGift *models.GiftPrize

	giftList := this.ServiceGift.GetAllUse()

	for _, gift := range giftList {

		// 中奖编码区间满足条件，说明可以中奖
		if gift.PrizeCodeA <= prizeCode && gift.PrizeCodeB >= prizeCode {
			if !limitBlack || gift.Gtype < conf.GiftTypeGiftSmall {
				prizeGift = &gift
				break
			}
		}
	}

	return prizeGift
}
