package routes

import (
	"github.com/kataras/iris/mvc"

	"../../bootstrap"
	"../../services"
	"../controllers"
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
}
