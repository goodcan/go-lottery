package controllers

import (
	"fmt"

	"github.com/kataras/iris"

	"go-lottery/comm"
	"go-lottery/models"
	"go-lottery/services"
)

type AdminBlackUserController struct {
	Ctx              iris.Context
	ServiceBlackUser services.BlackUserService
}

// GET /admin/blackUser
func (this *AdminBlackUserController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)
	page := this.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	// 数据列表
	dataList := this.ServiceBlackUser.GetAll()

	total := len(dataList)

	if len(dataList) >= size {
		total = int(this.ServiceBlackUser.CountAll())
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

// GET /admin/blackUser/black?id=1&time=0
func (this *AdminBlackUserController) GetBlack() {
	rs := comm.FromCtxGetResult(this.Ctx)

	id, err := this.Ctx.URLParamInt("id")

	t := this.Ctx.URLParamIntDefault("time", 0)

	if err == nil {
		if t > 0 {
			t = t*86400 + comm.NowUnix()
		}
		_ = this.ServiceBlackUser.Update(
			&models.BlackUser{
				Id:         id,
				BlackTime:  comm.StampToTime(t),
				SysUpdated: comm.NowTime(),
			},
			[]string{"black_time", "sys_updated"},
		)
	} else {
		rs.SetError(1, "missing black user id")
		this.Ctx.Next()
		return
	}

	this.Ctx.Next()
}
