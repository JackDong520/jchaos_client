package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func execNmap(stdout io.ReadCloser, error error) {

	reader := bufio.NewReader(stdout)
	firstLine := true
	outputText := ""
	print(error)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = escape(line)
		if strings.Contains(line, "open") {
			line = "<span class=\"open\">" + line + "</span>"
		} else if strings.Contains(line, "closed") {
			line = "<span class=\"closed\">" + line + "</span>"
		} else if strings.Contains(line, "filtered") {
			line = "<span class=\"filtered\">" + line + "</span>"
		}
		jump := "\n"
		if firstLine {
			jump = ""
		}
		outputText = outputText + jump + "<i>" + line + "</i>"

	}
	print(outputText)
	print("over")

}
func escape(str string) string {
	line := str
	strings.Replace(line, "&", "&amp;", -1)
	strings.Replace(line, "\"", "&quot;", -1)
	strings.Replace(line, "<", "&lt;", -1)
	strings.Replace(line, ">", "&gt;", -1)
	return line
}
func execute(nmapString string) {
	cmd := exec.Command("nmap", "-T4", "-F", nmapString)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	}
	_ = cmd.Start()
	go execNmap(stdout, err)
}

func main() {

	execute("10.59.13.137")
	var input string
	fmt.Scanln(&input)
}
