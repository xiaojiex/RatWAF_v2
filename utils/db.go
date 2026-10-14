package utils

import (
	"RatWAF_v2/dao"
	"RatWAF_v2/global"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func InsertDB(bite []byte) {
	data := dao.Waf{}
	err := json.Unmarshal(bite, &data)
	if err != nil {
		log.Printf("InsertDB格式化json错误%v", err)
		return
	}
	data.Createtime = time.Now()
	fmt.Println("攻击类型：", data.AttackType)
	err = global.DB.Create(&data).Error
	if err != nil {
		log.Printf("InsertDB插入数据错误%v", err)
		return
	}
	fmt.Println("插入数据库成功")
}
