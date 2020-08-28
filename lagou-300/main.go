package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/sillyhatxu/crawler-service/common/logconfig"
	"github.com/sillyhatxu/crawler-service/common/read"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	project  = "crawler-service"
	module   = "lagou-300"
	fileName = "test.html"
)

func init() {
	logconfig.InitialLogConfig(logconfig.Debug(true), logconfig.Env("dev"), logconfig.Project(project), logconfig.Module(module), logconfig.Version("v1.0.0-beta.1"))
}

func backups() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	logrus.Infof("read file : %s", filePathName)
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNode := htmlquery.FindOne(doc, "/html/body/div")
	titleXmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
	title := titleXmlElem.ChildText("/div")
	fmt.Println(title)
	xmlNode := htmlquery.FindOne(doc, "/html/body/div/div[4]/div")
	xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	for _, v := range xmlElem.ChildTexts("/p/span") {
		fmt.Println(v)
	}
	fmt.Println("------------")
	xmlElem = colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	for _, v := range xmlElem.ChildTexts("/h6/span") {
		fmt.Println(v)
	}
	fmt.Println("------------")
	xmlElem = colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	for _, v := range xmlElem.ChildAttrs("/p/img", "src") {
		fmt.Println(string(v))
	}
	fmt.Println("------------------------------------------------------------")
}

func main() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	logrus.Infof("read file : %s", filePathName)
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNode := htmlquery.FindOne(doc, "/html/body/div/div[4]/div")
	titleXmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)

	//
	//for _, v := range xmlElem.ChildAttrs("/p/img", "src") {
	//	fmt.Println(string(v))
	//}
	//ctx := &colly.Context{}
	//resp := &colly.Response{
	//	Request: &colly.Request{
	//		Ctx: ctx,
	//	},
	//	Ctx: ctx,
	//}
	//doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlPageByte))
	//if err != nil {
	//	panic(err)
	//}
	//i := 0
	//doc.Find("/html/body/div/div[4]/div").Each(func(_ int, s *goquery.Selection) {
	//	for _, n := range s.Nodes {
	//		htmlElement := colly.NewHTMLElementFromSelectionNode(resp, s, n, i)
	//		i++
	//		htmlElement.ForEach("", func(j int, element *colly.HTMLElement) {
	//			fmt.Println(element)
	//		})
	//	}
	//})
	//for _, child := range htmlquery.Find(tempXmlElem.DOM.(*html.Node), "/div") {
	//	fmt.Println(child)
	//	child.Attr
	//	//for _, attr := range child.Attr {
	//	//	if attr.Key == attrName {
	//	//		res = append(res, strings.TrimSpace(attr.Val))
	//	//	}
	//	//}
	//}
	//xmlNodes := htmlquery.Find(doc, "/html/body/div/div[4]/div")
	//for _, node := range xmlNodes {
	//	xmlElem = colly.NewXMLElementFromHTMLNode(resp, node)
	//	for _, v := range xmlElem.ChildAttr("/p/img","src") {
	//		fmt.Println(v)
	//	}
	//}
	//for _, xmlNode := range xmlNodes {
	//	fmt.Println(xmlNode)
	//	fmt.Println("----------")
	//	xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	//	value := xmlElem.ChildText("/p/span")
	//	if value != "" {
	//		fmt.Println(value)
	//	}
	//	value = xmlElem.ChildText("/h6/span")
	//	if value != "" {
	//		fmt.Println(value)
	//	}
	//}
}
