package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"public_utils/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/get-ip", &controllers.PublicApiController{}, "*:GetIP")
	beego.Router("/get-time", &controllers.PublicApiController{}, "*:GetTime")
}
