package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/sillyhatxu/crawler-service/common/logconfig"
	"github.com/sillyhatxu/crawler-service/common/read"
	"os"
	"strings"
)

const (
	project = "crawler-service"
	module  = "algoexpert-go/SystemsDesignFundamentals"
)

func init() {
	logconfig.InitialLogConfig(logconfig.Debug(true), logconfig.Env("dev"), logconfig.Project(project), logconfig.Module(module), logconfig.Version("v1.0.0-beta.1"))
}

func replace(input string) string {
	input = strings.ReplaceAll(input, "  ", "")
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "  ", "\n")
	return input
}

func main() {
	pwd, _ := os.Getwd()
	filePathName := fmt.Sprintf("%s/%s/%s.html", pwd, module, "test")
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))

	contentXmlNodes := htmlquery.Find(doc, "/html/body/div/*")
	for i, contentXmlNode := range contentXmlNodes {
		xmlElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
		title := xmlElem.ChildText("/div[2]/div[2]/h2")
		fmt.Println(fmt.Sprintf("%d. %s", (i + 1), title))

		content := xmlElem.ChildText("/div[2]/div[2]/p")
		fmt.Println(fmt.Sprintf("%s", replace(content)))

		prerequisites := xmlElem.ChildText("/div[3]/div/div")
		fmt.Println(fmt.Sprintf("%s", replace(prerequisites)))

		con1 := xmlElem.ChildText("/div[3]/div/div[2]")
		fmt.Println(fmt.Sprintf("%s", replace(con1)))

		con2 := xmlElem.ChildText("/div[4]/div")
		fmt.Println(fmt.Sprintf("%s", replace(con2)))

		fmt.Println("\n")
	}
}
