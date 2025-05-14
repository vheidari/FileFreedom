package FileDownloader

import ("net/http"
 	"fmt"
 	"log"
 	"io"
 	"os"
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
	 
	nFileDownloader := FileDownloader {
		m3u8FileUrl : nM3u8FileUrl,
		encryptionKeyUrl : nEncryptionKeyUrl,
		tsFileUrls : nTsFileUrls,
	}

	return nFileDownloader
}





func (url *FileDownloader) downloadM3u8File () {

	if url.m3u8FileUrl.url == "" {
		 log.Fatal("M3u8FileUrl.url : can't be empty")
	}

	m3u8FileName := "m3u8.m3u8"


	res, err := http.Get(url.m3u8FileUrl.url)

	if err != nil {
		 log.Fatal("")		
	}


	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		 log.Fatalf("Response failed with status code : %d and\n body : %s\n", res.StatusCode, body)
	}



	createFile, err := os.Create(m3u8FileName)

	fmt.Println(string(body))
	
	createFile.WriteString(string(body))
	
	// os.WriteFile(m3u8FileName, string(body), 0666)
	
		
}



func (url *M3u8FileUrl) DownloadIt () {

		
}
