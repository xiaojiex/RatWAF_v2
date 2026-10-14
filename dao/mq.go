package dao

import (
	"RatWAF_v2/global"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMq() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	global.MQConn = conn
	global.CH = ch

	ExchangeName := "AttackLog"
	//声明(创建)交换机
	err = global.CH.ExchangeDeclare(ExchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("交换器声明失败", err)
	}
	_, err = global.CH.QueueDeclare("log", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("队列声明失败", err)
	}
	err = global.CH.QueueBind("log", "", ExchangeName, false, nil)
	if err != nil {
		log.Fatalln("交换器绑定失败", err)
	}
	fmt.Println("initmq执行成功")
}

// func ConnectMq() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	global.MQConn = conn
// 	global.CH = ch
// }
