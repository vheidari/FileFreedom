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


type M3u8Data struct {
	 encryptionMethod string
	 encryptionKeyFile string
	 encryptionKeyUrl string 
	 videoFiles []string 
}


type FileDownloader  struct {
 	m3u8FileUrl M3u8FileUrl
 	m3u8Data *M3u8Data
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

	nDownloadPath  := DownloadPath.DownloadPath()
	 
	nFileDownloader := FileDownloader {
		m3u8FileUrl : nM3u8FileUrl,
		downloadPath: nDownloadPath,
	}

	return nFileDownloader
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



// M3u8 Manipulator Functions 

func MakeM3u8Data() *M3u8Data {

	var nM3u8Data M3u8Data =  M3u8Data {
		encryptionMethod : "",
		encryptionKeyFile : "",
		encryptionKeyUrl: "",
		videoFiles : make([]string, 0),
	}

	return &nM3u8Data
}

func (m3u8Data  *M3u8Data) SetM3u8EncryptionMethod(method string) {
		m3u8Data.encryptionMethod = method
}


func (m3u8Data  *M3u8Data) GetM3u8EncryptionMethod() string {
		return m3u8Data.encryptionMethod
}


func (m3u8Data  *M3u8Data) SetM3u8EncryptionKey(key string) {
	m3u8Data.encryptionKeyFile = key
}


func (m3u8Data  *M3u8Data) GetM3u8EncryptionKey() string {
	return m3u8Data.encryptionKeyFile
}


func (m3u8Data  *M3u8Data) SetM3u8VideoFiles(videoFiles *[]string) {
	m3u8Data.videoFiles = *videoFiles
}


func (m3u8Data  *M3u8Data) GetM3u8VideoFiles() *[]string {
	return &m3u8Data.videoFiles 
}


func (m3u8Data  *M3u8Data) SetM3u8EncryptionKeyUrl(url string) {
	m3u8Data.encryptionKeyUrl = url
}


func (m3u8Data  *M3u8Data) GetM3u8EncryptionKeyUrl() string {
	return m3u8Data.encryptionKeyUrl
}


func (m3u8Data  *M3u8Data) SetM3u8Data(nM3u8Data *M3u8Data) {
	m3u8Data = nM3u8Data
}


func (m3u8Data  *M3u8Data) GetM3u8Data() *M3u8Data {
	return	m3u8Data
}


func (url *M3u8FileUrl) DownloadIt () {

		
}
