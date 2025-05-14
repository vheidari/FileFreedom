package main


import "fmt"
import "net/http"
import "strings"
import "os"
import "log"

var Download string = "./Download"

var serverIp string = "127.0.0.1:8081"

func videoServer(w http.ResponseWriter, r *http.Request) {

	getUrl := r.URL.Path
	fmt.Println(getUrl)
	 	
	if 	getUrl != "" &&  getUrl != "/" {

		filePath := Download + getUrl
		
		fmt.Println(filePath)

		fileName := strings.Trim(getUrl, "/")
		
		w.Header().Set("Content-Disposition", "attachment; filename=" + fileName)
		
		http.ServeFile(w, r, filePath)		
	}

	ListOfFiles, err := os.ReadDir(Download)

	if err != nil {
		log.Fatal(err)
	}

	outputHtml := ""
	for _ , file := range ListOfFiles {
		outputHtml +=  "<li>http://" + serverIp + "/" + file.Name() + "</li>"	
	}

	titleText := "Video Stream Service - testServer 0/"
	fmt.Fprintln(w, "<html><head><title>" + titleText + "</title></head><body style='margin:auto;width:800px;border:1px solid green;text-align:center;margin-top:50px;border-radius:3px'><p style='border-bottom:1px solid green;color:green;'><b>" + titleText + "</b></p><div><p>List Of Files : </p><ul style='text-align: left;'>" + outputHtml  + "</ul></div></body></html>")
}


func main() {

	http.HandleFunc("/", videoServer)
	http.ListenAndServe(serverIp, nil)

}


