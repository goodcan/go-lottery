package controllers

import (
	"fmt"

	"github.com/kataras/iris"

	"go-lottery/comm"
	"go-lottery/models"
	"go-lottery/services"
)

type AdminResultController struct {
	Ctx           iris.Context
	ServiceResult services.ResultService
}

func (this *AdminResultController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)
	giftId := this.Ctx.URLParamIntDefault("gift_id", 0)
	uid := this.Ctx.URLParamIntDefault("uid", 0)
	page := this.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	var dataList []models.Result

	if giftId > 0 {
		dataList = this.ServiceResult.SearchByGift(giftId, page, size)
	} else if uid > 0 {
		dataList = this.ServiceResult.SearchByUser(uid, page, size)
	} else {
		dataList = this.ServiceResult.GetAll()
	}

	total := (page - 1) + len(dataList)

	// 超出每页大小再查询总数
	if len(dataList) >= size {

		if giftId > 0 {
			total = int(this.ServiceResult.CountByGift(giftId))
		} else if uid > 0 {
			total = int(this.ServiceResult.CountByUser(uid))
		} else {
			total = int(this.ServiceResult.CountAll())
		}

		pageNext = fmt.Sprintf("%d", page+1)
	}

	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	rs.Data = map[string]interface{}{
		"giftId":   giftId,
		"uid":      uid,
		"total":    total,
		"page":     page,
		"size":     size,
		"pagePrev": pagePrev,
		"pageNext": pageNext,
		"list":     dataList,
	}

	this.Ctx.Next()
}

func (this *AdminResultController) GetDelete() {
	rs := comm.FromCtxGetResult(this.Ctx)
	id, err := this.Ctx.URLParamInt("id")
	if err == nil || id == -1 {
		rs.SetError(1, "missing id")
		this.Ctx.Next()
		return
	}

	refer := this.Ctx.GetHeader("Referer")
	data := map[string]interface{}{}
	data["id"] = id
	if refer == "" {
		data["refer"] = "http://" + this.Ctx.Host() + "/admin/result"
	} else {
		data["refer"] = refer
	}

	rs.Data = data

	this.Ctx.Next()
}

func (this *AdminResultController) GetCheat() {
	rs := comm.FromCtxGetResult(this.Ctx)

	id, err := this.Ctx.URLParamInt("id")
	if err == nil || id == -1 {
		rs.SetError(1, "missing id")
		this.Ctx.Next()
		return
	}

	_ = this.ServiceResult.Update(
		&models.Result{Id: id, SysStatus: 2},
		[]string{"sys_status"},
	)

	refer := this.Ctx.GetHeader("Referer")
	data := map[string]interface{}{}
	data["id"] = id
	if refer == "" {
		data["refer"] = "http://" + this.Ctx.Host() + "/admin/result"
	} else {
		data["refer"] = refer
	}

	rs.Data = data

	this.Ctx.Next()
}

func (this *AdminResultController) GetReset() {
	rs := comm.FromCtxGetResult(this.Ctx)

	id, err := this.Ctx.URLParamInt("id")
	if err == nil || id == -1 {
		rs.SetError(1, "missing id")
		this.Ctx.Next()
		return
	}

	_ = this.ServiceResult.Update(
		&models.Result{Id: id, SysStatus: 0},
		[]string{"sys_status"},
	)

	refer := this.Ctx.GetHeader("Referer")
	data := map[string]interface{}{}
	data["id"] = id
	if refer == "" {
		data["refer"] = "http://" + this.Ctx.Host() + "/admin/result"
	} else {
		data["refer"] = refer
	}

	rs.Data = data

	this.Ctx.Next()
}
