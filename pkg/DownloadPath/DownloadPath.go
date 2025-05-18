package DownloadPath

import "os"
import "log"
import "errors"
import "runtime"

const (
	Linux string = "~/Download"
	Windows string = ""
)


// Retrieve the download path base on the operating system
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

	// Create the download dirctory if it is not exist
	if !isDirExist(downloadDir) {
		err := os.Mkdir(downloadDir, 0750)
			if err != nil   {
				log.Fatal(err)
		}
	}
	

	return downloadDir
}


// CreateDownloadPath takes two string input arguments: `path` and `fileName`.
// It ensures that the specified download path exists; if it does not, it creates the directory.
// The function returns the full download path as a string, which is a combination of the path and the fileName.
func CreateDownloadPath(path string, fileName string) string {

	// Ensure path and fileName aren't an empty string 
	if path == "" || fileName == "" {
		errors.New("paht or fileName can't be empty")
	}

	downloadPath := path  + "/" + fileName


	// Check if the Download Path isn't exist create one
	downloadPathDir := isDirExist(path)
	if !downloadPathDir {
		err := os.Mkdir(path, 0750)
		if err != nil   {
			log.Fatal(err)
		}
	}
	

	return downloadPath 
}


// CreateBaseDownloadPath takes one string input argument: `path`.
// It ensures that the specified download path exists; if it does not, it creates the directory.
// The function returns true if the specified path exists or if it was created successfully. 
func CreateBaseDownloadPath(path string) bool {

	if path == ""  {
		errors.New("path can't be an empty string")
	}

	// Check if the Download Path isn't exist create one
	downloadPathDir := isDirExist(path)
	if !downloadPathDir {
		err := os.Mkdir(path, 0750)
		if err != nil   {
			return false
		}
	}
	
	return true
}

// isDirExist function checks whether the specified directory path exists.
// The function returns true if directory exists; otherwise, it return false.
func isDirExist(path string) bool {

	_, err := os.Stat(path)

	if err != nil {
		return false 
	}

	return true
}
