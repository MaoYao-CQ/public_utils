package config

import (
	beego "github.com/beego/beego/v2/server/web"
)

var (
	Debug           bool
	Lang            string
	ProductName     string
	PublicApiDomain string
	IpdbPath        string
)

func Init() {
	//flag.StringVar(&RunType, "t", "", "运行模式")
	//flag.StringVar(&ListenHttpPort, "port", "15900", "监听端口")
	//flag.Parse()

	var err error
	PublicApiDomain, err = beego.AppConfig.String("PublicApiDomain")
	if nil != err || "" == PublicApiDomain {
		PublicApiDomain = "pub-utils.ai996.top"
	}

	ProductName, err = beego.AppConfig.String("ProductName")
	if nil != err || "" == ProductName {
		ProductName = "PUB-UTILS"
	}

	Lang, err = beego.AppConfig.String("Lang")
	if nil != err || "" == Lang {
		Lang = "zh-CN"
	}

	IpdbPath, err = beego.AppConfig.String("IpdbPath")
	if nil != err || "" == IpdbPath {
		IpdbPath = "conf/system/ipIpFree.ipdb"
	}

	Debug, err = beego.AppConfig.Bool("Debug")
	if nil != err {
		Debug = false
	}
}
