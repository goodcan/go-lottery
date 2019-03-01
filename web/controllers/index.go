package controllers

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"

	"../../comm"
	"../../conf"
	"../../models"
	"../../services"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}

// http://localhost:8080/
func (this *IndexController) Get() {
	rs := comm.FromCtxGetResult(this.Ctx)
	uri := "http://" + this.Ctx.Host()
	rs.Data = map[string]interface{}{
		"login":         uri + "/login",
		"logout":        uri + "/logout",
		"gifts":         uri + "/gifts",
		"loginUserList": uri + "/online",
		"admin":         uri + "/admin",
	}

	this.Ctx.Next()
}

// http://localhost:8080/gifts
func (this *IndexController) GetGifts() {
	rs := comm.FromCtxGetResult(this.Ctx)

	dataList := this.ServiceGift.GetAll()
	gifts := make([]models.Gift, 0)

	for _, data := range dataList {
		if data.SysStatus == 0 {
			gifts = append(gifts, data)
		}
	}

	rs.Data = map[string]interface{}{
		"gifts": gifts,
	}

	this.Ctx.Next()
}

// http://localhost:8080/new/prize
func (this *IndexController) GetNewPrize() {
	// TODO

}

// http://locahost:8080/login
func (this *IndexController) GetLogin() {
	uid := comm.RandInt(100000)
	loginUser := models.LoginUser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowTime(),
		Ip:       comm.ClientIp(this.Ctx.Request()),
	}

	loginUser.Sign = comm.CreateLoginUserSign(&loginUser)

	comm.SetLoginUser(this.Ctx.ResponseWriter(), &loginUser)
	//comm.Redirect(this.Ctx.ResponseWriter(),
	//	"/public/index.html?from=login")

	rs := comm.FromCtxGetResult(this.Ctx)
	rs.Data = loginUser

	conf.LoginUser.Store(loginUser.Uid, loginUser)

	this.Ctx.Next()
}

// http://locahost:8080/logout
func (this *IndexController) GetLogout() {
	comm.SetLoginUser(this.Ctx.ResponseWriter(), nil)
	//comm.Redirect(this.Ctx.ResponseWriter(),
	//	"/public/index.html?from=logout")

	rs := comm.FromCtxGetResult(this.Ctx)

	uid := this.Ctx.URLParamIntDefault("uid", 0)

	if uid == 0 {
		rs.SetError(1, "missing uid")
		this.Ctx.Next()
	}

	if _, ok := conf.LoginUser.Load(uid); ok {
		conf.LoginUser.Delete(uid)
	} else {
		rs.SetError(1, "用户不存在")
	}

	this.Ctx.Next()
}

// http://localhost:8080/online
func (this *IndexController) GetOnline() {
	rs := comm.FromCtxGetResult(this.Ctx)

	userList := make(map[string]interface{})
	conf.LoginUser.Range(func(key, value interface{}) bool {
		userList[strconv.Itoa(key.(int))] = value
		return true
	})

	rs.Data = userList

	this.Ctx.Next()
}
