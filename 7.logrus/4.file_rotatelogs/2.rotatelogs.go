package main

import(
	"log"
	"github.com/lestrrat-go/file-rotatelogs"
)

func main() {
	//rl, _ := rotatelogs.New("/tmp/accesslog2.txt.%Y%m%d%H%M")

	rl, _ := rotatelogs.New(
		"/tmp/log.%Y%m%d",
		rotatelogs.WithClock(rotatelogs.UTC),
	)

	log.SetOutput(rl)

	log.Printf("Hello, World!")
}
