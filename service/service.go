package service

import (
	"github.com/tidwall/gjson"
	"kunpeng/fileutil"
	"kunpeng/http"
)

func GetNmapInfoFromIp(ip string) {
	data := "nmapString=-T4 -F " + ip
	jsonString := http.HttpPost("http://127.0.0.1:8080/runnmap", data, "application/x-www-form-urlencoded")

	print(jsonString)
	value := gjson.Get(jsonString, "FileName")
	println(value.String())
	temppath := fileutil.GetCurrentDirectory()
	if fileutil.CheckIfFileExist(temppath + "/windows2012.txt") {
		print(temppath + "/windows2012.txt:" + "file exist")
	} else {
		fileutil.CreateFile(temppath + "/windows2012.txt")
		fileutil.AppendToFile(temppath+"/windows2012.txt", "["+jsonString+"]")
	}
	fileutil.AppendToFileFromSubOne(temppath+"/windows2012.txt", ","+jsonString+"]")
}
