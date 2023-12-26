package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"public_utils/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Get")
	beego.Router("/google965a213a981ca817.html", &controllers.HomeController{}, "*:Google")

	initRouter()
	beego.Router("/get-ip", &controllers.PublicApiController{}, "*:GetIpString")
	beego.Router("/get-time", &controllers.PublicApiController{}, "*:GetTime")

	//beego.InsertFilter("/api/*", beego.BeforeExec, filterCheckToken)

}

func initRouter() {

	nsPubApi := beego.NewNamespace("/api/v1",
		beego.NSNamespace("ip",
			//beego.NSInclude(&public_api.FaceEnController{}),
			beego.NSRouter("/get-string", &controllers.PublicApiController{}, "*:GetIpString"),
			beego.NSRouter("/get", &controllers.PublicApiController{}, "*:GetIpData"),
		),
		beego.NSNamespace("time",
			beego.NSRouter("/", &controllers.PublicApiController{}, "*:GetTime"),
		),
	)

	beego.AddNamespace(nsPubApi)

}
