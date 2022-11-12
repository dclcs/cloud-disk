# 初始化项目


password   ATBB7PLVFunSUu5nmg2jcA6Vv9bE92931EB9 / ATBBqHfhm5Ggc4v2UAjedRXWknYCA4C4D7A4



# 安装GO环境 
- 参考：[连接](https://learnku.com/articles/69929)

## 项目初始化
- `go mod init cloud-disk`
- `go get xorm.io/xorm`
- 安装超时问题： `go env -w GOPROXY=https://goproxy.cn,direct`


## 安装 go-zero
- `go install github.com/zeromicro/go-zero/tools/goctl@latest`
- 切到单体服务
    - `goctl api new core`
    - `goctl api go -api core.api -dir . -style go_zero`
    - 运行： `go run core.go -f etc/core-api.yaml`
- 安装验证码服务库
    - `go get github.com/jordan-wright/email`