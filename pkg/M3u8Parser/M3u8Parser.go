package M3u8Parser

import "vheidari/FileFreedom/pkg/FileDownloader"
import "os"
import "log"
import "strings"
import "bufio"



func ParseIt(downloadPath string, fileName string) (*FileDownloader.M3u8Data, error)  {

	// Genearting File Path
	dirPath := downloadPath
	filePath := dirPath + "/" + fileName


	// Opening M3u8 File 	
	fileData, err :=  os.Open(filePath)	

	if err != nil {
		log.Fatal(err)
	}

	// Free Resource When Our Jobs is Done
	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)
	var nM3u8Data = FileDownloader.MakeM3u8Data()
	videoFiles := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// Extract Encryption Method and Key
		if strings.HasPrefix(line, "#EXT-X-KEY:") {
			methodAndKey := strings.Split(line, ",")

			for _, item := range methodAndKey   {
				if strings.HasPrefix(item , "#EXT-X-KEY:METHOD=") {					
					nM3u8Data.SetM3u8EncryptionMethod( strings.TrimPrefix(item, "#EXT-X-KEY:METHOD=") )
				} else if strings.HasPrefix(item, "URI=") {
					if nM3u8Data != nil {
						nM3u8Data.SetM3u8EncryptionKey( strings.Trim(strings.TrimPrefix(item, "URI="), "\"") )
					}
				}
				
			}
				
		} else if !strings.HasPrefix(line, "#") {
			// Excract Video Segments
			videoFiles = append(videoFiles, line)
		}

	}
	
	// Handling error
	if err := scanner.Err(); err != nil {
		return nil, err;
	}

	// Set Video Files 
	nM3u8Data.SetM3u8VideoFiles(&videoFiles)

	return nM3u8Data, nil
}


