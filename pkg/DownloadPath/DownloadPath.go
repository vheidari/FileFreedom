package DownloadPath

import "os"
import "log"
import "runtime"

const (
	Linux string = "~/Download"
	Windows string = ""
)


func DownloadPath() string  {

	downloadDir := ""
	
	userHomeDirectory, err := os.UserHomeDir()

	// Can we get user home directory
	if err != nil {
		log.Fatal(err)
	}

	// Get os name :
	//	 - ( go tool dist list )
	osName := runtime.GOOS

	// Make download directory
	if osName == "linux" || osName == "freebsd" || osName == "openbsd" || osName == "netbsd" || osName == "darwin" || osName == "dragonfly"  {
		downloadDir += userHomeDirectory + "/Downloads/FileFreedom"	
	} else if osName == "windows" {
		downloadDir += userHomeDirectory + "\\Downloads\\FileFreedom"
	}


	return downloadDir
}
