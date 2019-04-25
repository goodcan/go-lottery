package bootstrap

import (
	"log"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"go-lottery/comm"
	"go-lottery/conf"
)

const (
	StaticAssets = "./public/"
	Favicon      = "favicon.ico"
)

type Configurator func(bootstrapper *Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnData time.Time
}

func New(appName, appOwner string, cfgList ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnData: comm.NowTime(),
	}

	for _, cfg := range cfgList {
		cfg(b)
	}

	return b
}

func (this *Bootstrapper) Bootstrap() *Bootstrapper {
	//this.SetupViews("./views")
	this.SetupErrorHandler()

	this.Favicon(StaticAssets + Favicon)
	//this.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	this.setupCron()

	this.Use(recover.New())
	this.Use(logger.New())

	return this
}

func (this *Bootstrapper) Listen(addr string, cfgList ...iris.Configurator) {
	err := this.Run(iris.Addr(addr), cfgList...)

	if err != nil {
		log.Fatal("bootstrap.Listen error ", err)
	}
}

func (this *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	//htmlEngine := iris.HTML(viewDir, ".html")

	// production 环境设置 false
	htmlEngine.Reload(true)

	htmlEngine.AddFunc("FromUnixTimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeFormShort)
	})

	htmlEngine.AddFunc("FromUnixTime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeForm)
	})

	this.RegisterView(htmlEngine)
}

func (this *Bootstrapper) SetupErrorHandler() {
	this.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     this.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.JSON(err)

		//ctx.ViewData("Err", err)
		//ctx.ViewData("Title", "Error")
		//ctx.View("shared/error.html")
	})
}

func (this *Bootstrapper) Configure(cfgList ...Configurator) {
	for _, cfg := range cfgList {
		cfg(this)
	}
}

func (this *Bootstrapper) setupCron() {
	// TODO
}
