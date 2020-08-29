package txt

import (
	"bufio"
	"bytes"
	"os"
)

type File struct {
	FilePath    string
	FileName    string
	FileContent bytes.Buffer
}

func (file *File) pathName() string {
	return file.FilePath + file.FileName
}

func (file *File) Write() error {
	if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
		err := os.Mkdir(file.FilePath, 0777)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(file.pathName())
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(file.FileContent.String())
	if err != nil {
		return err
	}
	return w.Flush()
}
