package main


func testProducer()  {
	body := []byte("ybx-body")
	ProduceMessage(TopicName, Key, body)
}

func main() {
	testProducer()
}
