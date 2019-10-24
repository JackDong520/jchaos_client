package service

import (
	"github.com/tiagorlampert/CHAOS/src/util"
	"github.com/tidwall/gjson"
	"kunpeng/http"
	"kunpeng/fileutil"
)

func getNmapInfoFromIp(ip string) {
	data := "nmapString=-T4 -F " + ip
	jsonString := http.HttpPost("http://127.0.0.1:8080/runnmap", data, "application/x-www-form-urlencoded")
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
