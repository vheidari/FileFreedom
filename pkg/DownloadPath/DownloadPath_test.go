package DownloadPath

import "testing"
import "os"
import "runtime"
import "log"



func TestDownloadPath(t *testing.T) {

	//
	getDownloadPath := DownloadPath()

	if getDownloadPath == "" {
		t.Errorf("getDownloadPath can't be empty")
	}

	// 
	osName := runtime.GOOS
	
	userHomeAddress, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	defaultDownloadPath := ""

	if osName == "windows" {
		defaultDownloadPath += userHomeAddress  + "\\Downloads\\FileFreedom"		
	} else {
		defaultDownloadPath += userHomeAddress  + "/Downloads/FileFreedom"
	}


	if defaultDownloadPath != getDownloadPath {
		t.Errorf("defaultDownloadPath : %q , getDownloadPath : %q ", defaultDownloadPath , getDownloadPath)
	}
	

	
	
}
