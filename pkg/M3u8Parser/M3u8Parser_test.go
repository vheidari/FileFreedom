package M3u8Parser

import "testing"
import "vheidari/FileFreedom/pkg/FileDownloader"


func TestParsIt(t *testing.T) {

	m3u8FileUrl := "http://127.0.0.1:8081/index-f5-v1-a1.m3u8"
	m3u8FileName := "master.m3u8"
	getFileDownloader := FileDownloader.GenerateFileDownloader(m3u8FileUrl)

	ParseIt(&getFileDownloader, m3u8FileName)


}
