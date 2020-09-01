package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strings"
	"time"
)

const (
	url = "https://youdao.com/w"
)

func search(word string) (string, error) {
	result := ""
	c := colly.NewCollector(
		colly.AllowedDomains("youdao.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)
	//在请求之前调用
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//如果在请求期间发生错误则调用
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	//收到回复后调用
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		//fmt.Println("Visited", string(r.Body))
	})

	c.OnHTML("div[class=trans-container] ul", func(e *colly.HTMLElement) {
		if result == "" {
			result = strings.ReplaceAll(e.Text, " ", "")
		}
	})
	url := fmt.Sprintf("%s/%s/", url, word)
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {
	file, err := os.Open("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/you-dao/word.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			words = append(words, word)
		}
	}
	var buffer bytes.Buffer
	for _, word := range words {
		result, err := search(word)
		if err != nil {
			fmt.Println(word, err)
		}
		buffer.WriteString(word + "|||" + result + "###")
		time.Sleep(2 * time.Second)
	}
	fmt.Println(buffer.String())
}
