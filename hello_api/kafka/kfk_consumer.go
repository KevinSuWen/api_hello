package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/astaxie/beego/logs"
)

type CumConsumer interface {
	StartConsumer(c CumConsumer, ctx context.Context)
	ProcessData(msg *sarama.ConsumerMessage) error
}

type KfkConsumer struct {
	CumConsumer
	name    string
	brokers []string
	topics  []string
}

func (ths *KfkConsumer) StartConsumer(c CumConsumer, ctx context.Context) {
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Consumer.Return.Errors = true
    logs.Info("brokers:%s; topics:%s;  kafka %s brokers ",ths.brokers,ths.topics,ths.name)
	consumer, err := cluster.NewConsumer(ths.brokers, "cum-server", ths.topics, config)
	if err != nil {
		logs.Error("create kafka consumer failed! topics:%s; error:%s",ths.topics,err.Error())
	}
	defer consumer.Close()
	go func() {
		for err := range consumer.Errors() {
			logs.Error("consumer topic %v error", ths.topics, err.Error())
		}
	}()
    logs.Info("connet successful!")
	for {
		select {
		case part,ok := <-consumer.Partitions():
			if !ok {
				logs.Info("cannot consumer data from topic %s", ths.topics)
				return
			}
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					logs.Debug("consumer from kafka topic %s partition %d", msg.Topic, msg.Partition)
					for retryTimes := 0; retryTimes < 5; retryTimes++ {
						err := c.ProcessData(msg)
						if err == nil{
							consumer.MarkOffset(msg, "")
							break
						}
					}
				}
			}(part)
		case <-ctx.Done():
			logs.Info("shutdown %s consumer.", ths.name)
			return
		}
	}

}
