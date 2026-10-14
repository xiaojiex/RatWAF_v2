package utils

import (
	"RatWAF_v2/global"
	"bytes"
	"context"
	"log"
)

func IfExist(ip string, AttackType string) {
	ctx := context.Background()
	buf := new(bytes.Buffer)
	buf.WriteString("Attack:")
	buf.WriteString(ip)
	buf.WriteString(":")
	buf.WriteString(AttackType)
	str, err := global.RDB.Get(ctx, buf.String()).Result()
	if err != nil {
		log.Println("utils ifexist get err:", err)
	}
	if str == "" {
		AttackTimes(ip, AttackType, buf.String())
	} else {
		global.RDB.Incr(ctx, buf.String())
	}
}

func AttackTimes(ip, AttackType, BufStr string) {
	//attack:ip:AttackType
	_, err := global.RDB.Set(context.Background(), BufStr, 1, 0).Result()
	if err != nil {
		log.Println("attacktime set err:", err)
	}
}
