package main

import (
	"os"
	"strings"
)

func main()  {

	shellCmd := "#!/bin/bash-0\r"
	shellCmd += "\r\n"
	shellCmd += "#!/bin/${{CommitID}}\r"
	shellCmd += "\r\n"

	f, _ := os.Create("./1.txt")

	shellCmd = strings.Replace(shellCmd, "\r", "", -1)
	_, _ = f.WriteString(shellCmd)

	_ = f.Chmod(0755)
	_ = f.Close()
}