package fileutil

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func CheckIfFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func CreateFile(file_name string) {
	//创建文件
	f, err := os.Create(file_name)
	//判断是否出错
	if err != nil {
		fmt.Println(err)
	}
	//打印文件名称
	fmt.Println(f.Name())
	defer f.Close()
}

/**
从倒数第二位的位置进行写入
*/
func AppendToFileFromSubOne(file, str string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0660)

	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	off, err := f.Seek(-1, io.SeekEnd)
	if err != nil {
		fmt.Println("Seek err:", err)
		return
	}
	_, err = f.WriteAt([]byte(str), off)
	if err != nil {
		fmt.Println("WriteAt err:", err)
		return
	}
	fmt.Println("write successful")
	fmt.Println("off:", off)
}

func AppendToFile(file, str string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	f.WriteString(str)
}
func main() {
	filepath := "C:/Users/DZKD/AppData/Local/Temp/2019-10-22-20-04-30_windwos.txt"
	CreateFile(filepath)
}
