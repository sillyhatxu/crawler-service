package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	filedownload "github.com/sillyhatxu/crawler-service/common/download"
	"github.com/sillyhatxu/crawler-service/common/logconfig"
	"github.com/sillyhatxu/crawler-service/common/read"
	"github.com/sillyhatxu/crawler-service/common/txt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	project  = "crawler-service"
	module   = "algoexpert-go"
	fileName = "test.html"
)

var subtitles = []string{
	"一、", "二、", "三、", "四、", "五、",
	"六、", "七、", "八、", "九、", "十、",
	"十一、", "十二、", "十三、", "十四、", "十五、",
	"十六、", "十七、", "十八、", "十九、", "二十、",
	"二十一、", "二十二、", "二十三、", "二十四、", "二十五、",
	"二十六、", "二十七、", "二十八、", "二十九、", "三十、",
}

func init() {
	logconfig.InitialLogConfig(logconfig.Debug(true), logconfig.Env("dev"), logconfig.Project(project), logconfig.Module(module), logconfig.Version("v1.0.0-beta.1"))
}

func algoExpert() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	logrus.Infof("read file : %s", filePathName)
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/h2")
	titleXmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
	title := strings.Trim(strings.ReplaceAll(titleXmlElem.Text, "\n", ""), " ")
	fmt.Println(title)

	difficultyXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/div[1]/div[1]/span[2]")
	difficultyXmlElem := colly.NewXMLElementFromHTMLNode(resp, difficultyXmlNode)
	fmt.Println("difficulty :", difficultyXmlElem.Attr("data-tip"))

	categoryXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/div[1]/div[2]/span[2]")
	categoryXmlElem := colly.NewXMLElementFromHTMLNode(resp, categoryXmlNode)
	fmt.Println("Category :", categoryXmlElem.Text)

	contentXmlNodes := htmlquery.Find(doc, "/html/body/div/div/div/div[2]/div[2]/div[2]/*")
	for _, contentXmlNode := range contentXmlNodes {
		contentXmlElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
		content := strings.ReplaceAll(contentXmlElem.Text, "  ", "")
		fmt.Println(content)
	}
	//contentXmlElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
	//content := strings.Trim(strings.ReplaceAll(contentXmlElem.Text, "\n", " "), "  ")
	//content := strings.ReplaceAll(strings.ReplaceAll(contentXmlElem.Text, "\n", " "), "  ", "")
	//content := strings.ReplaceAll(contentXmlElem.Text, "  ", "")
	//fmt.Println(content)
	//xmlNode := htmlquery.FindOne(doc, "/html/body/div/div[4]/div")
	//xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	//for _, v := range xmlElem.ChildTexts("/p/span") {
	//	fmt.Println(v)
	//}
	//fmt.Println("------------")
	//xmlElem = colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	//for _, v := range xmlElem.ChildTexts("/h6/span") {
	//	fmt.Println(v)
	//}
	//fmt.Println("------------")
	//xmlElem = colly.NewXMLElementFromHTMLNode(resp, xmlNode)
	//for _, v := range xmlElem.ChildAttrs("/p/img", "src") {
	//	fmt.Println(string(v))
	//}
	//fmt.Println("------------------------------------------------------------")
}

func download(fileName string) {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	logrus.Infof("read file : %s", filePathName)
	var buffer bytes.Buffer
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNode := htmlquery.FindOne(doc, "/html/body/div")
	titleXmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
	title := titleXmlElem.ChildText("/div")
	buffer.WriteString(title + "\n" + "\n")
	xmlNodes := htmlquery.Find(doc, "/html/body/div/div[4]/div/*")
	subtitleIndex := 0
	downloadPath := fmt.Sprintf("%s/%s/%s/", pwd, module, title)
	imageIndex := 1
	for _, xmlNode := range xmlNodes {
		xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)
		switch xmlElem.Name {
		case "p":
			value := xmlElem.ChildText("/span")
			if value != "" {
				buffer.WriteString(value + "\n")
				continue
			}
			value = xmlElem.ChildAttr("/img", "src")
			if value != "" {
				filename, err := filedownload.BuildFileName(value)
				if err != nil {
					panic(err)
				}
				file := filedownload.File{
					FileURL:      value,
					DownloadPath: downloadPath,
					DownloadName: fmt.Sprintf("%d.%s", imageIndex, filename),
				}
				//fmt.Println(file)
				err = file.Download()
				if err != nil {
					panic(err)
				}
				buffer.WriteString(fmt.Sprintf("%d-%s \n", imageIndex, value))
				imageIndex++
				continue
			}
			buffer.WriteString("\n")
		case "h6":
			value := xmlElem.ChildText("/span")
			if value != "" {
				buffer.WriteString(subtitles[subtitleIndex] + value + "\n")
				subtitleIndex++
			} else {
				buffer.WriteString("\n")
			}
		case "h2":
			buffer.WriteString("\n")
		case "ul":
			for i, v := range xmlElem.ChildTexts("/li/p/span") {
				buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, v))
			}
		default:
			buffer.WriteString("\n")
		}
	}
	file := txt.File{
		FilePath:    downloadPath,
		FileName:    fmt.Sprintf("%s.txt", title),
		FileContent: buffer,
	}
	err := file.Write()
	if err != nil {
		panic(err)
	}
}

func main() {
	algoExpert()
}
