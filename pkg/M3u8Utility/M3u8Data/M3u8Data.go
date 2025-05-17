package M3u8Data


type M3u8Data struct {
	 encryptionMethod string
	 encryptionKeyFile string
	 url string 
	 videoFiles []string 
}


// M3u8 Manipulator Functions 
func MakeM3u8Data() *M3u8Data {

	var nM3u8Data M3u8Data =  M3u8Data {
		encryptionMethod : "",
		encryptionKeyFile : "",
		url: "",
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


func (m3u8Data  *M3u8Data) SetM3u8Url(url string) {
	m3u8Data.url = url
}


func (m3u8Data  *M3u8Data) GetM3u8Url() string {
	return m3u8Data.url
}


func (m3u8Data  *M3u8Data) SetM3u8Data(nM3u8Data *M3u8Data) {
	*m3u8Data = *nM3u8Data
}


func (m3u8Data  *M3u8Data) GetM3u8Data() *M3u8Data {
	return	m3u8Data
}

