package M3u8Parser

import "testing"
import "vheidari/FileFreedom/pkg/DownloadPath"
import "fmt"

func TestParsIt(t *testing.T) {

	// Get .m3u8 File Location
	getDownloadPath := DownloadPath.DownloadPath()
	m3u8FileName := "master.m3u8"

	// Parse .m3u8 File
	GetM3u8Data, err := ParseIt(getDownloadPath, m3u8FileName)

	if err != nil {
		 t.Errorf("Can't Parse .m3u8 file, Path: %q , FileName:%q, Error:%q ", getDownloadPath, m3u8FileName, err)
	}

	fmt.Println(GetM3u8Data)
	
}
