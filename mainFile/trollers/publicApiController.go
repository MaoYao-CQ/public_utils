package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/ipipdotnet/ipdb-go"
	"public_utils/config"
	"public_utils/services"
	"public_utils/services/utilHttp"
	"strconv"
	"time"
)

type PublicApiController struct {
	beego.Controller
}

func (c *PublicApiController) GetIpString() {
	var isDebug = c.GetString("debug")

	r := c.Ctx.Request
	debug := false
	if "true" == isDebug {
		debug = true
	}
	var ip = utilHttp.ClientPublicIP(r, debug)
	if "" == ip {
		ip = utilHttp.ClientIP(r, debug)
	}

	c.Ctx.WriteString(ip)
}

func (c *PublicApiController) GetTime() {
	format := c.GetString("t")
	var writeString string

	switch format {
	case "unix":
		writeString = strconv.FormatInt(time.Now().Unix(), 10)
	case "f11":
		writeString = time.Now().Format("2006-01-02 15:04:05")
	case "f12":
		writeString = time.Now().Format("2006-01-02 15:04")
	case "f13":
		writeString = time.Now().Format("2006-01-02")
	case "f21":
		writeString = time.Now().Format("2006年01月02日 15时04分05秒")
	case "f22":
		writeString = time.Now().Format("2006年01月02日 15时04分")
	case "f23":
		writeString = time.Now().Format("2006年01月02日")
	case "f31":
		writeString = time.Now().Format("20060102150405")
	case "f32":
		writeString = time.Now().Format("200601021504")
	case "f33":
		writeString = time.Now().Format("20060102")
	default:
		writeString = strconv.FormatInt(time.Now().Unix(), 10)
	}

	c.Ctx.WriteString(writeString)
}

type IpData struct {
	Ip          string `json:"ip"`
	CityName    string `json:"city_name"`
	RegionName  string `json:"region_name"`
	CountryName string `json:"country_name"`
}

func (c *PublicApiController) GetIpData() {

	ApiRj := services.ApiReturnJson{Status: false, Code: 500, Message: i18n.Tr(config.Lang, "api_return_json_default_message")}
	defer func() {
		c.Data["json"] = ApiRj
		_ = c.ServeJSON()
	}()

	var ip = c.GetString("ip")

	if ip == "self" {
		r := c.Ctx.Request
		ip = utilHttp.ClientPublicIP(r, false)
		if "" == ip {
			ip = utilHttp.ClientIP(r)
		}
	}

	if "" == ip {
		ApiRj.Message = i18n.Tr(config.Lang, "public_api_get_ip_data_ip_not_fond")
		return
	}

	db, err := ipdb.NewCity(config.IpdbPath)
	if err != nil {
		ApiRj.Message = i18n.Tr(config.Lang, "public_api_get_ip_data_ipdb_error")
		ApiRj.Errors = append(ApiRj.Errors, err.Error())
		return
	}

	ipInfo, err := db.FindInfo(ip, "CN")
	if err != nil {
		ApiRj.Message = i18n.Tr(config.Lang, "public_api_get_ip_data_find_error")
		ApiRj.Errors = append(ApiRj.Errors, err.Error())
		return
	}

	var ipData IpData
	ipData.Ip = ip
	ipData.CityName = ipInfo.CityName
	ipData.RegionName = ipInfo.RegionName
	ipData.CountryName = ipInfo.CountryName

	ApiRj.Data = ipData
	ApiRj.Status = true
	ApiRj.Code = 200
	ApiRj.Message = i18n.Tr(config.Lang, "api_return_json_ok_message")
	return

}
