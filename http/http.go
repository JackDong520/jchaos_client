package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiagorlampert/CHAOS/src/util"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func Get(url string) string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
// content:请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	req.Header.Add(`content-type`, contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func HttpPost(url string, data string, contentType string) string {
	resp, err := http.Post(url,
		contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}


func main() {
	//url := "http://127.0.0.1:8080"
	//print(Get(url))
	//contentType := "application/x-www-form-urlencoded"

	data := "nmapString=-T4 -F 10.59.13.169"
	jsonString := HttpPost("http://127.0.0.1:8080/runnmap", data, "application/x-www-form-urlencoded")
	print(jsonString)
	value := gjson.Get(jsonString, "FileName")
	println(value.String())
	temppath := util.GetCurrentDirectory()
	if util.CheckIfFileExist(temppath + "/windows2012.txt") {
		print(temppath + "/windows2012.txt:" + "file exist")
	} else {
		util.CreateFile(temppath + "/windows2012.txt")
		util.AppendToFile(temppath+"/windows2012.txt", "["+jsonString+"]")
	}
	util.AppendToFileFromSubOne(temppath+"/windows2012.txt", ","+jsonString+"]")

}
