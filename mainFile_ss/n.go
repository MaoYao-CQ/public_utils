package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"public_utils/config"
	"public_utils/controllers/common"
	_ "public_utils/routers"
	"public_utils/services/utilsLogger"
	"strings"
)
type langType struct {
	Lang string
	Name string
}
func init()  {
	langTypesStr, _ := beego.AppConfig.String("lang::types")
	langNamesStr, _ := beego.AppConfig.String("lang::names")
	languages := strings.Split(langTypesStr, "|")
	names := strings.Split(langNamesStr, "|")
	langTypes := make([]*langType, 0, len(languages))
	for i, v := range languages {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range languages {
		if err := i18n.SetMessage(lang, "conf/language/"+"locale_"+lang+".ini"); err != nil {
			fmt.Println(err)
			return
		}
	}

	config.Init()

	utilsLogger.Init()

	beego.ErrorController(&common.ErrorController{})

}

func main() {
	beego.Run()
}
