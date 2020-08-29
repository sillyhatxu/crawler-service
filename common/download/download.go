package download

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type File struct {
	FileURL      string
	DownloadPath string
	DownloadName string
}

func (file *File) Download() (err error) {
	err = file.buildFilePath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(file.DownloadPath); os.IsNotExist(err) {
		err := os.Mkdir(file.DownloadPath, 0777)
		if err != nil {
			return err
		}
	}
	// Build fileName from file url
	if file.DownloadName == "" {
		file.DownloadName, err = BuildFileName(file.FileURL)
		if err != nil {
			return err
		}
	}
	emptyFile, err := os.Create(file.filePathName())
	if err != nil {
		return err
	}
	// copy content on file
	return file.copyFile(emptyFile, file.httpClient())
}

func (file *File) filePathName() string {
	return file.DownloadPath + file.DownloadName
}

func (file *File) copyFile(emptyFile *os.File, client *http.Client) error {
	resp, err := client.Get(file.FileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//_, err = io.Copy(emptyFile, resp.Body)
	size, err := io.Copy(emptyFile, resp.Body)
	if err != nil {
		return err
	}
	defer emptyFile.Close()
	logrus.Infof("Just Downloaded a file %s with size %d", file.filePathName(), size)
	return nil
}

func (file *File) buildFilePath() error {
	if file.DownloadPath != "" {
		return nil
	}
	pwd, _ := os.Getwd()
	file.DownloadPath = pwd
	return nil
}

func BuildFileName(fileURL string) (string, error) {
	fileUrl, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1], nil
}

func (file *File) httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return &client
}
