package dao

import (
	"RatWAF_v2/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Waf struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	AttackerIP string    `gorm:"column:attacker_ip; not null" json:"attacker_ip"` // 攻击者ip
	Method     string    `gorm:"column:method; not null" json:"method"`
	Body       string    `gorm:"column:body; not null" json:"body"` // 攻击请求体
	Createtime time.Time `gorm:"column:createtime;autoCreateTime" json:"createtime"`
	Remote     string    `gorm:"column:remote; not null" json:"remote"`           // 攻击路径remote
	AttackType string    `gorm:"column:attack_type; not null" json:"attack_type"` // 攻击类型
}

func InitMySql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/waf?charset=utf8mb4&parseTime=True&loc=Local"
	logLevel := logger.Info
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}

	global.DB = db
	err = global.DB.AutoMigrate(&Waf{})
	if err != nil {
		log.Fatalf("自动迁移建表失败:%v", err)
	}
	fmt.Println("mysql初始化成功")
}

//后面还有根据redis内容插入表
