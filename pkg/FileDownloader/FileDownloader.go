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

	// Check if the input argument m3u8Url is an empty string.
	if m3u8Url == "" {
	 	return nil ,errors.New("m3u8Url cant be empty")
	}

 	// Create a new instance of M3u8FileUrl with the specified URL.
	nM3u8FileUrl := M3u8FileUrl {
	 	url: m3u8Url,
	}

	// Create a new instance of M3u8Data.
	nM3u8Data := M3u8Data.MakeM3u8Data()

	// Retrieve the download path from the DownloadPath instance.
	nDownloadPath  := DownloadPath.DownloadPath()
	 
	nFileDownloader := &FileDownloader {
		m3u8FileUrl : nM3u8FileUrl,
		m3u8Data: nM3u8Data,
		downloadPath: nDownloadPath,
	}

	return nFileDownloader, nil
}


// downloadM3u8File download `.m3u8` file with the specified `m3u8FileUrl.url` in the FileDownloader struct.
func (url *FileDownloader) downloadM3u8File () {

	if url.m3u8FileUrl.url == "" {
		 log.Fatal("M3u8FileUrl.url : can't be empty")
	}

	// Creates the download path directory (if it does not exist) and returns the full download path for a .m3u8 file.
	getDownloadPath := url.downloadPath
	m3u8FileName :=  "master.m3u8"
	fullDownloadPath := DownloadPath.CreateDownloadPath(getDownloadPath, m3u8FileName)

	// Download `.m3u8` file from specified video URL 
	res, err := http.Get(url.m3u8FileUrl.url) // Send a GET request to download the `.m3u8` file

	if err != nil {
		 log.Fatal("Faild to download .m3u8 file: %v", err)		
	}

	// Read the body content from the response and stores it in the `body` variable
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	
	if res.StatusCode > 299 {
		 log.Fatalf("Response failed with status code : %d and\n body : %s\n", res.StatusCode, body)
	}


	// Creates a new file on disk at the specified path
	createFile, err := os.Create(fullDownloadPath)

	if err != nil {
		 log.Fatalf("Problem to create %s file on the Disk : %s", m3u8FileName,  err)
	}
	
	createFile.WriteString(string(body))
			
}


// DownloadPath returns the download path as a string associated with the FileDownloader.
func (fileDownloader *FileDownloader) DownloadPath() string {
	return fileDownloader.downloadPath
}



func (url *M3u8FileUrl) DownloadKeyAndVideoFiles () {
	
}

// GetM3u8FileUrl returns a pointer to the current M3u8FileUrl associated with the FileDownloader.
func (fileDownloader *FileDownloader) GetM3u8FileUrl() *M3u8FileUrl {
	return &fileDownloader.m3u8FileUrl
}



// GateBaseUrl takes a name of the m3u8file and returns a Baseurl as string    
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
