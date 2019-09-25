package main

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const expectedMsg = "This is the expected test message."
const unexpectedMsg = "This message should not be logged."

func main() {
	log := logrus.New()

	tmpfile, err := ioutil.TempFile("/tmp", "456.txt")
	if err != nil {
		panic("error"+err.Error())
	}

	fname := tmpfile.Name()
	defer func() {
		tmpfile.Close()
		os.Remove(fname)
	}()

	hook := lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel: fname,
		logrus.WarnLevel: fname,
	}, nil)
	log.Hooks.Add(hook)

	log.Info(expectedMsg)
	log.Warn(unexpectedMsg)

	contents, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		panic("Error while reading from tmpfile:"+ err.Error())
	}

	fmt.Println("contents:%s, unexpectedMsg:%s, fname:%s", contents, unexpectedMsg, fname)
}