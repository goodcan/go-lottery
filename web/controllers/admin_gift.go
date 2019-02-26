package controllers

import (
	"github.com/kataras/iris"

	"../../services"
)

type AdminGiftController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}
