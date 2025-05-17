package M3u8Data

import "testing"
import "reflect"


// M3u8 Manipulator Functions 
func TestMakeM3u8Data(t *testing.T) {

	 expectM3u8Data := &M3u8Data {
	 	encryptionMethod: "",
	 	encryptionKeyFile: "",
	 	url: "",
	 	videoFiles: make([]string, 0),
	 }
	 nM3u8Data :=  MakeM3u8Data()


     // reflect.DeepEqual compare each filed in objects
	 if !reflect.DeepEqual(expectM3u8Data,  nM3u8Data) {
	 	t.Errorf("expectM3u8Data:%q , nM3u8Data:%q ", expectM3u8Data, nM3u8Data)
	 }

	 
}

func TestSet_Get_M3u8EncryptionMethod(t *testing.T) {
	expectMethod := "AES-128"
	nM3u8Data := MakeM3u8Data()
	nM3u8Data.SetM3u8EncryptionMethod(expectMethod)

	if expectMethod != nM3u8Data.GetM3u8EncryptionMethod() {
		t.Errorf("expectMethod:%q, nM3u8Data.GetEncryptionMethod:%q", expectMethod, nM3u8Data.GetM3u8EncryptionMethod() )
	}
	
		
}

func TestSet_Get_M3u8EncryptionKey(t *testing.T) {
	expectKey := "encryption-f5.key"
	nM3u8Data := MakeM3u8Data()
	nM3u8Data.SetM3u8EncryptionKey(expectKey)


	if expectKey != nM3u8Data.GetM3u8EncryptionKey() {
		t.Errorf("expectKey:%q, nM3u8Data.GetEncryptionKey:%q", expectKey, nM3u8Data.GetM3u8EncryptionKey() )
	}
	
}


func TestSet_Get_M3u8VideoFiles(t *testing.T) {
	expectVideos := make([]string, 0)

	expectVideos = append(expectVideos,"seg-1-f5-v1-a1.ts")
	expectVideos = append(expectVideos,"seg-2-f5-v1-a1.ts")

	nM3u8Data := MakeM3u8Data()
	nM3u8Data.SetM3u8VideoFiles(&expectVideos)

	resultVideos := nM3u8Data.GetM3u8VideoFiles()


	// reflect.DeepEqual compare each filed in objects
	if !reflect.DeepEqual( &expectVideos, resultVideos) {
		t.Errorf("expectVideos:%q, resultVideos:%q", expectVideos, resultVideos )
	}
}



func TestSet_Get_M3u8EncryptionKeyUrl(t *testing.T) {
	expectM3u8Url := "http://127.0.0.1/media"

	m3u8Data := MakeM3u8Data()
	m3u8Data.SetM3u8Url(expectM3u8Url)
	resultUrl := m3u8Data.GetM3u8Url()

	if expectM3u8Url != resultUrl {
		t.Errorf("expectM3u8Url:%q, resultUrl:%q", expectM3u8Url, resultUrl )		
	}
	
}



func TestSet_Get_M3u8Data(t *testing.T) {
	expectM3u8Data := M3u8Data {
		encryptionMethod: "AES-127",
		encryptionKeyFile : "encryption-f5.key",
		url: "http://127.0.0.1/media",
		videoFiles: make([]string, 0),
	}

	m3u8Data := MakeM3u8Data()
	m3u8Data.SetM3u8Data(&expectM3u8Data)
	resultM3u8Data := m3u8Data.GetM3u8Data()

	if !reflect.DeepEqual( &expectM3u8Data, resultM3u8Data ){
		t.Errorf("expectM3u8Url:%q, resultUrl:%q", expectM3u8Data, resultM3u8Data )		
	}
}
