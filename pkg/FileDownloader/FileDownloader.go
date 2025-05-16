package FileDownloader

import ("net/http"
 	"log"
 	"io"
 	"os"
 	"vheidari/FileFreedom/pkg/DownloadPath"
 )

type M3u8FileUrl struct {
	url string
		
}

type EncryptionKeyUrl struct {
	url string 
}


type TsFileUrls struct {
	urls map[uint]string	
}

type FileDownloader  struct {
 	m3u8FileUrl M3u8FileUrl
 	encryptionKeyUrl EncryptionKeyUrl
 	tsFileUrls TsFileUrls
 	downloadPath string
}


func GenerateFileDownloader(m3u8Url string) FileDownloader {

	// check input arrgument that can't be a empty url
	if m3u8Url == "" {
	 	log.Fatal("m3u8Url cant be empty") 	
	}


 	 // generate a M3u8FileUrl object
	nM3u8FileUrl := M3u8FileUrl {
	 	url: m3u8Url,
	}

	nEncryptionKeyUrl := EncryptionKeyUrl {
	 	url : "",
	}

	nTsFileUrls := TsFileUrls {
		urls : make(map[uint]string),
	}

	nDownloadPath  := DownloadPath.DownloadPath()
	 
	nFileDownloader := FileDownloader {
		m3u8FileUrl : nM3u8FileUrl,
		encryptionKeyUrl : nEncryptionKeyUrl,
		tsFileUrls : nTsFileUrls,
		downloadPath: nDownloadPath,
	}

	return nFileDownloader
}





func (url *FileDownloader) downloadM3u8File () {

	if url.m3u8FileUrl.url == "" {
		 log.Fatal("M3u8FileUrl.url : can't be empty")
	}


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
