package common

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"public_utils/config"
	"public_utils/models"
)

type ErrorController struct {
	beego.Controller
}
func InitError(c *ErrorController, code int, message string){
	if c.IsAjax() {
		c.Data["json"] = models.ApiReturnJson{Status: false, Code: code, Message: message}
		_ = c.ServeJSON()
	} else {
		c.Data["code"] = code
		c.Data["title"] = config.ProductName
		c.Data["message"] = message
		c.TplName = "error/layout.tpl"
	}
}

func (c *ErrorController) Error404() {
	InitError(c, 404, i18n.Tr(config.Lang, "page_not_found"))
}

func (c *ErrorController) Error403() {
	InitError(c, 403, i18n.Tr(config.Lang, "page_not_found"))
}

func (c *ErrorController) Error401() {
	InitError(c, 401, i18n.Tr(config.Lang, "unauthenticated"))
}
