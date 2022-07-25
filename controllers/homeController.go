package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"public_utils/config"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "UTILS|工具喵"
	if config.Debug {
		c.Data["Website"] = "DEBUG-工具喵"
	}

	c.Data["timestamp"] = time.Now().Unix()
	c.Data["datetime"] = time.Now().Format("2006-01-02 15:04:05")

	c.Data["Email"] = "xxxxxx@163.com"
	c.TplName = "index.tpl"
}
