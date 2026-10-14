package utils

import (
	"RatWAF_v2/regex"
)

func CheckReq(req1, req2 string) (bool, string) {
	if isAttack, attackType := regex.SqlInjection(req1, req2); isAttack { //bool str
		return true, attackType
	}
	if isAttack, attackType := regex.RCE(req1, req2); isAttack {
		return true, attackType
	}
	if isAttack, attackType := regex.XSS(req1, req2); isAttack {
		return true, attackType
	}
	return false, ""
}

/*
这是开始的redis模拟队列，现采用mq
func GetRpop() string {
	ctx := context.Background()
	for {
		// 设置 5 秒超时
		result, err := global.RDB.BRPop(ctx, 5*time.Second, "http_attack_log").Result()
		if err == redis.Nil {
			//5秒内没数据，继续等待
			continue
		} else if err != nil {
			// 情况 2：Redis 发生真正的错误（如断网），打印日志并稍作休眠
			log.Printf("GetRpop 错误: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		// result[1] JSON 字符串
		jsonData := result[1]
		InsertDB(jsonData)
	}
}
*/
