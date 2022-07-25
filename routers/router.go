package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"public_utils/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	initRouter()

	//beego.InsertFilter("/api/*", beego.BeforeExec, filterCheckToken)

}

func initRouter() {

	nsPubApi := beego.NewNamespace("/api/v1",
		beego.NSRouter("/get-ip", &controllers.PublicApiController{}, "*:GetIpString"),
		beego.NSRouter("/get-time", &controllers.PublicApiController{}, "*:GetTime"),
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
