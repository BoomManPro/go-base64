package main

import (
	"encoding/base64"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	//编译命令
	//go build -ldflags="-H windowsgui"

	filePath:=os.Args[1]
	bytes, _ := ioutil.ReadFile(filePath)
	encodeString := base64.StdEncoding.EncodeToString(bytes)
	ioutil.WriteFile(path.Base(filePath)+".base64.txt",[]byte(encodeString),os.ModePerm)
	clipboard.WriteAll(encodeString)

}
