package controllers

import (
	"fmt"

	"github.com/kataras/iris"

	"go-lottery/comm"
	"go-lottery/models"
	"go-lottery/services"
)

type AdminBlackIpController struct {
	Ctx            iris.Context
	ServiceBlackIp services.BlackIpService
}

// GET /admin/blackIp
func (this *AdminBlackIpController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)
	page := this.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	// 数据列表
	dataList := this.ServiceBlackIp.GetAll()

	total := len(dataList)

	if len(dataList) >= size {
		total = int(this.ServiceBlackIp.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	rs.Data = map[string]interface{}{
		"total":    total,
		"page":     page,
		"size":     size,
		"pagePrev": pagePrev,
		"pageNext": pageNext,
		"list":     dataList,
	}

	this.Ctx.Next()
}

// GET /admin/blackIp/black?id=1&time=0
func (this *AdminBlackIpController) GetBlack() {
	rs := comm.FromCtxGetResult(this.Ctx)

	id, err := this.Ctx.URLParamInt("id")

	t := this.Ctx.URLParamIntDefault("time", 0)

	if err == nil {
		if t > 0 {
			t = t*86400 + comm.NowUnix()
		}
		_ = this.ServiceBlackIp.Update(
			&models.BlackIp{
				Id:         id,
				BlackTime:  comm.StampToTime(t),
				SysUpdated: comm.NowTime(),
			},
			[]string{"black_time", "sys_updated"},
		)
	} else {
		rs.SetError(1, "missing black ip id")
		this.Ctx.Next()
		return
	}

	this.Ctx.Next()
}
