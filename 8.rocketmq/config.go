package main

var NameServers []string
var ProducerGroupName string
var TopicName string
var Key	string

var ConsumerGroup string
var BrokerNames   []string
var QueueSize     int



func init() {
	NameServers = append(NameServers, "127.0.0.1:9876")
	ProducerGroupName = "ProducerGroupName-ybx"
	TopicName = "ybx-topic1"
	Key = "ybx-key"

	ConsumerGroup = "ConsumerGroup-ybx"
}