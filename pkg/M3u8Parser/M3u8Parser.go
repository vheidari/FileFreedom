package M3u8Parser

import "vheidari/FileFreedom/pkg/FileDownloader"
import "os"
import "log"
import "strings"
import "bufio"

type m3u8Data struct {
	 encryptionMethod string
	 encryptionKeyFile string
	 encryptoinKeyUrl string 
	 videoFiles []string 
}


func ParseIt(fileDownloader *FileDownloader.FileDownloader, fileName string) error {

	// Genearting File Path
	dirPath := fileDownloader.DownloadPath()
	filePath := dirPath + "/" + fileName


	// Opening M3u8 File 	
	fileData, err :=  os.Open(filePath)	

	if err != nil {
		log.Fatal(err)
	}

	// Free Resource When Our Jobs Was Done
	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)
	var nM3u8Data  *m3u8Data

	for scanner.Scan() {
		line := scanner.Text()

		// Extract Encryption Method and Key
		if strings.HasPrefix(line, "#EXT-X-KEY:") {
			methodAndKey := strings.Split(line, ",")

			for _, item := range methodAndKey   {
				if strings.HasPrefix(item , "#EXT-X-KEY:METHOD=") {
					nM3u8Data = &m3u8Data { encryptionMethod: strings.TrimPrefix(item, "#EXT-X-KEY:METHOD=")}
				} else if strings.HasPrefix(item, "URI=") {
					if nM3u8Data != nil {
						nM3u8Data.encryptionKeyFile = strings.Trim(strings.TrimPrefix(item, "URI="), "\"")
					}
				}
				
			}
				
		} else if !strings.HasPrefix(line, "#") {
			// Excract Video Segments
			nM3u8Data.videoFiles = append(nM3u8Data.videoFiles, line)
		}

	}
	
	// Handling error
	if err := scanner.Err(); err != nil {
		return err;
	}


	
	return nil
}


