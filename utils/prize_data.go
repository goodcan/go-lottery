package utils

import (
	"log"

	"go-lottery/comm"
	"go-lottery/services"
)

func PrizeGift(id int, giftService services.GiftService) bool {
	rows, err := giftService.DecrLeftNum(id, 1)
	if rows < 1 || err != nil {
		log.Println(
			"prize_data.PrizeGift giftService.DecrLeftNum error=", err,
			"rows=", rows)
		return false
	}
	return true
}

func PrizeCodeDiff(id int, codeService services.CodeService) string {
	lockUid := 0 - id - 100000000

	ok := LockLucky(lockUid)
	if !ok {
		return ""
	} else {
		defer UnLockLucky(lockUid)
	}

	codeId := 0

	codeInfo := codeService.NextUsingCode(id, codeId)

	if codeInfo != nil && codeInfo.Id > 0 {
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowTime()
		codeService.Update(codeInfo, nil)
		return codeInfo.Code
	} else {
		log.Println("prize_data.PrizeCodeDiff num codeInfo, gift_id=", id)
		return ""
	}
}
