package DownloadPath

import "testing"
import "os"
import "runtime"
import "log"



// TestDownloadPath verifies that the DownloadPath function works correctly.
func TestDownloadPath(t *testing.T) {

	// Retrieve the download path using the DownloadPath function.
	getDownloadPath := DownloadPath()

	// Ensure that the retrieved download path is not an empty string.
	if getDownloadPath == "" {
		t.Errorf("getDownloadPath can't be empty")
	}

	// Get the oprationg system name.
	osName := runtime.GOOS

	// Get the user's home Directory.
	userHomeAddress, err := os.UserHomeDir()

	// Checks if there is error during.
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the defautl download path.
	defaultDownloadPath := ""

	// Construct the default download path base on the operating syste.
	if osName == "windows" {
		defaultDownloadPath += userHomeAddress  + "\\Downloads\\FileFreedom"		
	} else {
		defaultDownloadPath += userHomeAddress  + "/Downloads/FileFreedom"
	}

	// Compare the default download path with the retrieved download path.
	if defaultDownloadPath != getDownloadPath {
		t.Errorf("defaultDownloadPath : %q , getDownloadPath : %q ", defaultDownloadPath , getDownloadPath)
	}
	

	
}


// TestCreateDownloadPath verifies the CreateDownloadPath function work correctly.
// It ensure that the CreateDownloadPath can return a full path string of downloadPhat+fileName
func TestCreateDownloadPath(t *testing.T)  {

	fileName := "master.m3u8"
	downloadPath := DownloadPath()
	expectFullDownloadPath := downloadPath  + "/" + fileName

	fullDownloadPath := CreateDownloadPath(downloadPath, fileName)


	if expectFullDownloadPath != fullDownloadPath {
		t.Errorf("expectFullDownloadPath:%q, fullDownloadPath:%q", expectFullDownloadPath, fullDownloadPath)
	}
}

// TestCreateBaseDownloadPath verifies that the CreateBaseDownloadPath can create a base download path.
// even if specified download path does not exist 
func TestCreateBaseDownloadPath(t *testing.T)  {

	downloadPath := DownloadPath()

	isCreated := CreateBaseDownloadPath(downloadPath)

	if !isCreated {
		t.Errorf("Faild to create directory : %q", downloadPath) 
	}

}


// TestIsDirExist verifies that the isDirExist function correctly checks whether the specified download path exist or does not exist.
func TestIsDirExist(t *testing.T) {

	downloadPath := DownloadPath()

	isCreated := CreateBaseDownloadPath(downloadPath)
		
	isDirectoryExist := isDirExist(downloadPath)

	if !isCreated || !isDirectoryExist {
		 t.Errorf("The download Path directory does not exist!")
	}
	
			
}
