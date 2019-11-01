package nmap

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

var tmpFileName = os.TempDir()

type NmapInfo struct {
	Ip       string
	PathName string
}

var NmapInfos = make(map[string]NmapInfo)

func execnewNmap(stdout io.ReadCloser, error error, ip string) {

	reader := bufio.NewReader(stdout)
	print(error)
	filename := "\\Nmap-scan_" + time.Now().Format("2006-01-02_15-04-05_") + ip + ".txt"

	objectfilname := tmpFileName + filename
	nmapinfo := NmapInfo{Ip: ip, PathName: objectfilname}

	NmapInfos[ip] = nmapinfo

	f, err := os.Create(objectfilname)
	if err != nil {
		fmt.Println(err)
		return
	}
	print("create new file :" + objectfilname + "\n")
	/**

	 */
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}

		l, err := f.WriteString(line)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		print(line)
		fmt.Println(l, "bytes written successfully")

	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//print(outputText)
	//print("over")

}
func escape(str string) string {
	line := str
	strings.Replace(line, "&", "&amp;", -1)
	strings.Replace(line, "\"", "&quot;", -1)
	strings.Replace(line, "<", "&lt;", -1)
	strings.Replace(line, ">", "&gt;", -1)
	return line
}
func newexecute(ip string) {
	cmd := exec.Command("nmap", "-T4", "-F", ip)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	}
	_ = cmd.Start()
	go execnewNmap(stdout, err, ip)
}
func ExecNmap() {
	newexecute("10.59.13.137")
	newexecute("10.59.13.138")
	newexecute("10.59.13.136")
	newexecute("10.59.13.135")
	newexecute("10.59.13.134")
}
