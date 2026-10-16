package utils

import (
	"RatWaf_v2/global"
	"context"
	"fmt"
	"log"
	"time"
)

//	if GetNum(ip) > 200 {
//			err := SetBan(ip, 12*time.Hour)
//			if err != nil {
//				log.Println("set ban err:", err)
//			}
//		}
//
// incr并评估
// INCR 计数后判断是否触发封禁，返回本 IP当前是否已处于封禁态
func IncrAndEvaluate(ip string) (banned bool, err error) {
	key := fmt.Sprintf("TTL:%s", ip)
	if GetNum(ip) == 0 {
		Setdata(ip)
	} else { //增加次数
		global.RDB.Incr(context.Background(), key)
		if GetNum(ip) > 200 {
			err := SetBan(ip, 12*time.Hour)
			if err != nil {
				log.Println("set ban err:", err)
			}
		}
	}
	ok, err := IsBanned(ip)
	return ok, err
}

// 记录攻击次数
func Setdata(ip string) {
	key := fmt.Sprintf("TTL:%s", ip)
	global.RDB.Set(context.Background(), key, 1, time.Minute)
}

// 获取攻击次数
func GetNum(ip string) int {
	key := fmt.Sprintf("TTL:%s", ip)
	res, err := global.RDB.Get(context.Background(), key).Int()
	if err != nil {
		log.Println(err)
	}
	return res
}

// 检测链路调用，判断某 IP 是否被封
func IsBanned(ip string) (bool, error) {
	key := fmt.Sprintf("ban:%s", ip)
	n, err := global.RDB.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// 写入封禁键并设 TTL（重复调用刷新 TTL）
func SetBan(ip string, duration time.Duration) error {
	key := fmt.Sprintf("ban:%s", ip)
	_, err := global.RDB.Set(context.Background(), key, true, duration).Result()
	return err
}

///	存在固定窗口边界漏洞

// 14:40:30 首次攻击 → 窗口 [14:40:30, 14:41:30]，你打 120 次（<200，不封）
// 14:41:30 key 过期归零 → 新窗口 [14:41:30, 14:42:30]，再打 120 次（仍 <200）

// func IncrAndEvaluate1(ip string) (bool, error) {
// 	ctx := context.Background()
// 	key := fmt.Sprintf("TTL:%s", ip)

// 	cnt, err := global.RDB.Incr(ctx, key).Result()
// 	if err != nil {
// 		return false, err
// 	}
// 	if cnt == 1 { // 只在首次创建时设窗口过期
// 		if err := global.RDB.Expire(ctx, key, time.Minute).Err(); err != nil {
// 			return false, err
// 		}
// 	}
// 	if cnt >= 200 {
// 		if err := SetBan(ip, 12*time.Hour); err != nil {
// 			return false, err
// 		}
// 		return true, nil
// 	}
// 	return false, nil
// }

//
//方案 B（真·滑动窗口，面试能讲透）​：用 ZSET 记每次请求的时间戳，只数"最近 60 秒"内的条数，没有边界漏洞：
// time.Now().UnixMilli()

// global.RDB.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: member})   // member 需唯一(时间戳+随机数)
// global.RDB.ZRemRangeByScore(ctx, key, "0", fmt.Sprint(now-60000))         // 删 60s 外旧记录
// cnt, _ := global.RDB.ZCard(ctx, key).Result()                             // 窗口内数量
// global.RDB.Expire(ctx, key, time.Minute).Err()
// return cnt >= threshold, nil
