package utils

import (
	"RatWAF_v2/global"
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Creator(Attack string) {
	//发送消息
	err := global.CH.PublishWithContext(
		context.Background(),
		"AttackLog",
		"",
		true,
		false,
		amqp.Publishing{
			Body: []byte(Attack),
		},
	)
	if err != nil {
		log.Println("消息创建失败utils/Creator", err)
		return
	}
	fmt.Println("消息存入成功")
}
func Consumer() {
	fmt.Println("consumer 已启动")
	q, err := global.CH.QueueInspect("log")
	if err != nil {
		log.Fatal(err)
		return
	}
	consume, err := global.CH.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for msg := range consume {
		// fmt.Println("-------------------------------")
		fmt.Println(string(msg.Body))
		//从mq中取出直接存入mysql
		InsertDB(msg.Body)
		// fmt.Println("收到消息")
	}
}
