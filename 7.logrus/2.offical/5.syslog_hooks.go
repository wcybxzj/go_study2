package main

import (
  "github.com/sirupsen/logrus"
  lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
  log       := logrus.New()
  hook, err := lSyslog.NewSyslogHook("", "", "/tmp/123.txt", "")

  if err == nil {
    log.Hooks.Add(hook)
  }
}
