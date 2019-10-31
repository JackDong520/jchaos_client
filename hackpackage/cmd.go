package hackpackage

import (
	"io/ioutil"
	"os/exec"
)

func ExecCmd(msg string) []byte {
	cmd := exec.Command("cmd", msg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		print("err1")
		print(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		print("err1")
		print(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		print(err)
	}
	return opBytes
}
