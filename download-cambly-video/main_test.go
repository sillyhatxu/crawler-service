package main

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	format, err := time.Parse("Jan 2, 2006 3:04 PM", "Jun 7, 2019 8:32 PM")
	//format, err := time.Parse("Jan 02 06 15:04 MST", "Jun 7, 2019 8:32 PM")
	if err != nil {
		t.Fatalf("string to time format error: %v", err)
	}
	fmt.Println(format.Format("2006-01-02"))

}

func TestFormatTimeNow(t *testing.T) {
	fmt.Println(time.Now().Format("Jan 02, 2006 03:04 PM"))
}

func TestDownloadFile(t *testing.T) {
	url := redirect("www.example.com", "")
	downloadFile("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/test.mp4", url)
}
