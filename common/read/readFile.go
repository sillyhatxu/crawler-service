package read

import (
	"io/ioutil"
	"os"
)

func File(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	contextByte, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return contextByte
}
