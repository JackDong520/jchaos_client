package main

import (
	"io/ioutil"
	"os/exec"
)

func main() {

	cmd := exec.Command("cmd", "")
	//stdin, _ := cmd.StdinPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		print("err1")
		print(err)
	}
	defer stdout.Close()
	//mystdin := cmd.Stdout

	if err := cmd.Start(); err != nil {
		print("err1")
		print(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		print(err)
	}
	print(string(opBytes))
	cmd.Start()
	//buf := "cd d:"
	//
	//mystdin.Write([]byte(buf))
	//opBytess, err := ioutil.ReadAll(stdout)
	//if err != nil {
	//	print(err)
	//}
	//print(string(opBytess))

}
