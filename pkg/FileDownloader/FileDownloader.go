package FileDownloader

import ("net/http"
 	"log"
 	"io"
 	"os"
 	"errors"
 	"strings"
 	"vheidari/FileFreedom/pkg/DownloadPath"
 	"vheidari/FileFreedom/pkg/M3u8Utility/M3u8Data"
 )

type M3u8FileUrl struct {
	url string		
}



type FileDownloader  struct {
 	m3u8FileUrl M3u8FileUrl
 	m3u8Data *M3u8Data.M3u8Data
 	downloadPath string
}

func GenerateFileDownloader(m3u8Url string) (*FileDownloader, error) {

	// Is Imput Arggument An Empty String
	if m3u8Url == "" {
	 	return nil ,errors.New("m3u8Url cant be empty")
	}

 	 // Make A M3u8FileUrl Instance
	nM3u8FileUrl := M3u8FileUrl {
	 	url: m3u8Url,
	}

	// Make An Instance of the M3u8Data 
	nM3u8Data := M3u8Data.MakeM3u8Data()

	// Geting Download Path
	nDownloadPath  := DownloadPath.DownloadPath()
	 
	nFileDownloader := &FileDownloader {
		m3u8FileUrl : nM3u8FileUrl,
		m3u8Data: nM3u8Data,
		downloadPath: nDownloadPath,
	}

	return nFileDownloader, nil
}



func (url *FileDownloader) downloadM3u8File () {

	if url.m3u8FileUrl.url == "" {
		 log.Fatal("M3u8FileUrl.url : can't be empty")
	}

	// Creating Download Path
	getDownloadPath := url.downloadPath
	m3u8FileName :=  "master.m3u8"
	downloadPath := getDownloadPath + "/" + m3u8FileName

	// Get file from m3u8 file from Service
	res, err := http.Get(url.m3u8FileUrl.url)

	if err != nil {
		 log.Fatal("")		
	}

	// Read Body Content
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	
	if res.StatusCode > 299 {
		 log.Fatalf("Response failed with status code : %d and\n body : %s\n", res.StatusCode, body)
	}


	// Write file body on disk
	createFile, err := os.Create(downloadPath)

	if err != nil {
		 log.Fatalf("Problem to create %s file on the Disk : %s", m3u8FileName,  err)
	}
	
	createFile.WriteString(string(body))
			
}



func (fileDownloader *FileDownloader) DownloadPath() string {
	return fileDownloader.downloadPath
}



func (url *M3u8FileUrl) DownloadIt () {

		
}


func (fileDownloader *FileDownloader) GetM3u8FileUrl() *M3u8FileUrl {
	return &fileDownloader.m3u8FileUrl
}


func (url *M3u8FileUrl) GetBaseUrl(m3u8FileName string) string {
	var found bool
	var before string
	if strings.HasSuffix(url.url, m3u8FileName ) {
		before, found = strings.CutSuffix(string(url.url), m3u8FileName)
		if found {
			return before
		} 
	}
	return before
		
}
