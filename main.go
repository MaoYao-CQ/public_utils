package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "public_utils/routers"
)

func main() {
	beego.BConfig.AppName = "pub-utils"
	beego.BConfig.RunMode = "prod"
	beego.BConfig.Listen.HTTPPort = 11080
	beego.BConfig.Listen.HTTPSPort = 11443
	beego.BConfig.Listen.HTTPAddr = "pub-utils.local.mozhe.cn"
	beego.Run()
}

