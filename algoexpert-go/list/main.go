package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/sillyhatxu/crawler-service/common/read"
	"os"
	"strings"
)

const (
	project = "crawler-service"
	module  = "algoexpert-go/list"
)

func replace(input string) string {
	input = strings.ReplaceAll(input, "  ", "")
	input = strings.ReplaceAll(input, "\n", "")
	return input
}

func main() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s.html", pwd, module, "test")
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNodes := htmlquery.Find(doc, "/html/body/div/*")
	for i, titleXmlNode := range titleXmlNodes {
		xmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
		title := replace(xmlElem.ChildText("/h2"))
		if title == "" {
			continue
		}
		titleContent := strings.Split(title, "-")
		questionCount := strings.Split(titleContent[1], "/")
		//fmt.Println(fmt.Sprintf("%d. %s-%s(%s)", i+1, titleContent[0], questionCount[1], "Easy"))
		folder1 := fmt.Sprintf("%d. %s(%s)", i+1, titleContent[0], questionCount[1])
		fmt.Println(folder1)
		_ = os.Mkdir("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/test/"+folder1, 0755)
		contentXmlNodes := htmlquery.Find(titleXmlNode, "/div/*")
		index := 1
		for j, contentXmlNode := range contentXmlNodes {
			contentElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
			difficulty := "Unknown"
			if j%2 == 1 {
				diffXmlNode := htmlquery.FindOne(contentXmlNode, "/span")
				diffElem := colly.NewXMLElementFromHTMLNode(resp, diffXmlNode)
				if diffElem.Attr("class") == "_3eKEN1BLI2EEjKbMeiKrUY lHK2uX3P5XJdmXvmsIP4B" {
					difficulty = "Easy"
				} else if diffElem.Attr("class") == "_3eKEN1BLI2EEjKbMeiKrUY _1iSGcXmk6fsuzIU6WmB9cg" {
					difficulty = "Medium"
				} else if diffElem.Attr("class") == "_3eKEN1BLI2EEjKbMeiKrUY _21WHDydrGZohZgq-EYVGlb" {
					difficulty = "Hard"
				} else if diffElem.Attr("class") == "_3eKEN1BLI2EEjKbMeiKrUY _1b4MQdsE1wmXLHcxIju2PQ" {
					difficulty = "Very Hard"
				}
				subTitle := replace(contentElem.Text)
				folder2 := fmt.Sprintf("%d.%d %s(%s)", i+1, index, subTitle, difficulty)
				fmt.Println(folder2)
				_ = os.Mkdir("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/test/"+folder2, 0755)
				index++
			}
		}
	}
}
