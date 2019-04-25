package routes

import (
	"github.com/kataras/iris/mvc"

	"go-lottery/bootstrap"
	"go-lottery/services"
	"go-lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	userDayService := services.NewUserDayService()
	codeService := services.NewCodeService()
	giftService := services.NewGiftService()
	resultService := services.NewResultService()
	blackIpService := services.NewBlackIpService()
	blackUserService := services.NewBlackUserService()

	index := mvc.New(b.Party("/"))

	index.Register(
		userDayService,
		codeService,
		giftService,
		resultService,
		blackIpService,
		blackUserService,
	)

	index.Handle(new(controllers.IndexController))

	admin := index.Party("/admin")
	admin.Handle(new(controllers.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Handle(new(controllers.AdminGiftController))

	adminCode := admin.Party("/code")
	adminCode.Handle(new(controllers.AdminCodeController))

	adminResult := admin.Party("/result")
	adminResult.Handle(new(controllers.AdminResultController))

	adminBlackUser := admin.Party("/blackUser")
	adminBlackUser.Handle(new(controllers.AdminBlackUserController))

	adminBlackIp := admin.Party("/blackIp")
	adminBlackIp.Handle(new(controllers.AdminBlackIpController))
}
