# 公开工具项目
	更新时间:2021年11月12日
## 介绍
	这是一个公共工具项目、可以放在公网、不设权限访问；
	目前拥有工具
		1、获取客户端真实出网IP
## 编译命令
	windows
		bee pack -be GOOS=windows
	linux
		 bee pack -be GOOS=linux -be CGO_ENABLED=0
