package FileDownloader

import "testing"
import "fmt"



func TestGenerateFileDownloader(t *testing.T) {

	// 
	m3u8FileUrl := "https://127.0.0.1:8081/index-f5-v1-a1.m3u8"
	
	fileDownloaderObj, err := GenerateFileDownloader(m3u8FileUrl)

	if err != nil {
		t.Error(err)
	}

	if fileDownloaderObj.m3u8FileUrl.url !=  m3u8FileUrl {
		 // [todo] :
		 t.Error("m3u8Urls isn't match")
	}

}




func TestDownloadM3u8File(t *testing.T) {

	// 
	m3u8FileUrl := "http://127.0.0.1:8081/index-f5-v1-a1.m3u8"
	
	fileDownloaderObj, err := GenerateFileDownloader(m3u8FileUrl)

	if err != nil {
		t.Error(err)
	}

	// [todo] :
	fileDownloaderObj.downloadM3u8File()
		
}






func TestGetBaseUrl(t *testing.T) {

	// 
	m3u8FileUrl := "http://127.0.0.1:8081/index-f5-v1-a1.m3u8"
	
	fileDownloaderObj, err := GenerateFileDownloader(m3u8FileUrl)

	if err != nil {
		t.Error(err)
	}


	getM3u8FileUrl := fileDownloaderObj.GetM3u8FileUrl()
	getBaseUrl := getM3u8FileUrl.GetBaseUrl("index-f5-v1-a1.m3u8")
	fmt.Println(getBaseUrl)
	

	
	// [todo] :
	fileDownloaderObj.downloadM3u8File()
		
}






