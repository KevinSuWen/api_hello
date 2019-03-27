package kafka

import (
	"context"
	"strings"
	"cum/kafka/pb"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var ctx context.Context
var cancel context.CancelFunc

func Start() {
	logs.Info("Kafka is start1!")
}

func init() {
	kfkBasicBrokers :=  beego.AppConfig.String("kfkBasicBrokers")
	ctx, cancel = context.WithCancel(context.Background())
	logs.Info("kafka server start")
	basicConsumer := newBasicConsumer("base", strings.Split(kfkBasicBrokers, ","), []string{kafka100.TopicDeviceOnlineData})
	go basicConsumer.StartConsumer(basicConsumer, ctx)
}


