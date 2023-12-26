# 公开工具项目
	更新时间:2021年11月12日
## 介绍
	这是一个公共工具项目、可以放在公网、不设权限访问；
	目前拥有工具
        1、获取客户端真实出网IP
        2、生成口令
        3、IP归属查询
        4、时间戳与时间转换
## 编译命令
### # 本地调试
    bee run

### # windows平台部署
    bee pack -be GOOS=windows -be CGO_ENABLED=0 -exp="vue:.idea:.git" -exs=".go:.DS_Store:.tmp:go.mod:go.sum:.exe:.exe~"

### # linux平台部署
    bee pack -be GOOS=linux -be CGO_ENABLED=0 -exp="vue:.idea:.git" -exs=".go:.DS_Store:.tmp:go.mod:go.sum:.exe:.exe~"
