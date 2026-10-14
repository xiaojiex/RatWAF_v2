package main

import (
	"RatWAF_v2/dao"
	"RatWAF_v2/proxy"
	"RatWAF_v2/utils"
	"RatWAF_v2/web"
)

func main() {
	dao.InitRedis()
	dao.InitMySql()
	dao.InitMq()
	go utils.Consumer()
	//读取redis并写入mysql，这里是阻塞的所以用协程
	//go utils.GetRpop()
	go web.InitWeb()
	proxy.Reverse()
}
