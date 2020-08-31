package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/sillyhatxu/crawler-service/common/read"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
)

const (
	project  = "crawler-service"
	module   = "english-at-work"
	fileName = "test1.html"
)

func readFile() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	logrus.Infof("read file : %s", filePathName)
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	urlXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div[2]")
	urlXmlElem := colly.NewXMLElementFromHTMLNode(resp, urlXmlNode)
	fmt.Println("https://www.bbc.co.uk" + urlXmlElem.ChildAttr("/a", "href"))

	urlXmlNodes := htmlquery.Find(doc, "/html/body/div/div[2]/ul/*")
	for _, urlXmlNode := range urlXmlNodes {
		urlXmlElem = colly.NewXMLElementFromHTMLNode(resp, urlXmlNode)
		fmt.Println("https://www.bbc.co.uk" + urlXmlElem.ChildAttr("/div[2]/a", "href"))
	}
}

func search(url string) (string, error) {
	var buffer bytes.Buffer
	c := colly.NewCollector(
		colly.AllowedDomains("bbc.co.uk", "www.bbc.co.uk", "https://www.bbc.co.uk"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
		//colly.Debugger(&debug.LogDebugger{}),
	)
	c.SetRequestTimeout(30 * time.Second)
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
	//c.OnXML("/html/body/div[6]/div[3]/div/div", func(element *colly.XMLElement) {
	//	title := element.ChildText("/div[3]/h3")
	//	buffer.WriteString(fmt.Sprintf("#%s", title))
	//	context := element.ChildTexts("/div[6]/div/*")
	//	fmt.Println(context)
	//})
	c.OnHTML("div[id=orb-modules]", func(e *colly.HTMLElement) {
		//c.OnHTML("div[class=widget-container]", func(e *colly.HTMLElement) {
		//c.OnHTML("div[class=widget-container widget-container-left]", func(e *colly.HTMLElement) {
		//fmt.Println(e)
		e.ChildText(".widget widget-heading clear-left")
		e.ForEach(".widget-container-left", func(_ int, leftElement *colly.HTMLElement) {
			leftElement.ForEach(".widget .widget-heading .clear-left > h3", func(_ int, titleElement *colly.HTMLElement) {
				buffer.WriteString(fmt.Sprintf("#%s", titleElement.Text))
			})
			fmt.Println(leftElement)
		})
		buffer.WriteString("")
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	//readFile()
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, "url.txt")
	file, err := os.Open(filePathName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}
	result, err := search("https://www.bbc.co.uk/learningenglish/english/features/english-at-work/66-language-for-a-wedding-day")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	//for _, url := range urls {
	//	result, err := search(url)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}
