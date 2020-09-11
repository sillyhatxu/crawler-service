package main

import (
	"fmt"
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
	fmt.Println("title :", title)

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
}

func main() {
	algoExpert()
}
