package FileDownloader

import "testing"
// import "fmt"



func TestGenerateFileDownloader(t *testing.T) {

	// 
	m3u8FileUrl := "https://test.com/master.m3u8"
	
	fileDownloaderObj := GenerateFileDownloader(m3u8FileUrl)

	if fileDownloaderObj.m3u8FileUrl.url !=  m3u8FileUrl {
		 t.Error("m3u8 Url isn't match ")
	}
		
}




func TestDownloadM3u8File(t *testing.T) {

	// 
	m3u8FileUrl := "http://127.0.0.1:8081/index-f5-v1-a1.m3u8"
	
	fileDownloaderObj := GenerateFileDownloader(m3u8FileUrl)

	fileDownloaderObj.downloadM3u8File()
		
}



