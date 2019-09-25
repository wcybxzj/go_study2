package main

import (
	"context"
	"zuji/common/dlog"

	"github.com/apache/rocketmq-client-go"
	"github.com/apache/rocketmq-client-go/primitive"
	"github.com/apache/rocketmq-client-go/producer"
)

var prod rocketmq.Producer

func ProducerConnect() error {
	var err error

	prod, err = rocketmq.NewProducer(
		producer.WithNameServer(NameServers),
		producer.WithGroupName(ProducerGroupName),
		producer.WithRetry(2))

	if err != nil {
		prod = nil
		dlog.WriteLog("RocketMQ Connect error ", err)
		return err
	}

	err = prod.Start()
	if err != nil {
		prod = nil
		dlog.WriteLog("RocketMQ Start error ", err)
		return err
	}

	return nil
}

func ProducerClose() error {
	if prod != nil {
		return prod.Shutdown()
	}

	return nil
}

func ProduceMessage(topic string, key string, body []byte) error {
	if prod == nil {
		err := ProducerConnect()
		if err != nil {
			dlog.WriteLog("RocketMQ Connect error", err)
			return err
		}
	}

	res, err := prod.SendSync(context.Background() ,
							&primitive.Message{
								Topic:  topic,
								Body:   body,
								Properties: map[string]string{"KEYS": key},
							})
	if err != nil {
		prod = nil
		dlog.WriteLog("RocketMQ SendSync error", err)
		return err
	}

	if res.Status != primitive.SendOK {
		dlog.WriteLog("RocketMQ SendSync error", res)
	}

	return nil
}
