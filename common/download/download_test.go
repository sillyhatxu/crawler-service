package download

import (
	"testing"
)

func TestFile_Download(t *testing.T) {
	file := File{
		FileURL: "http://s0.lgstatic.com/i/image2/M01/9D/32/CgoB5l2td9qAf2VuAABerHKNze4736.png",
	}
	err := file.Download()
	if err != nil {
		t.Fatalf("download file error: %v", err)
	}
	t.Log("download file success")
}
