package controllers

import (
	"github.com/kataras/iris"

	"go-lottery/services"
)

type AdminGiftController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}
