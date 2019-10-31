package hackpackage

import (
	"github.com/matishsiao/goInfo"
	"strconv"
)

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
