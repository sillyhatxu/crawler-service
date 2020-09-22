package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/sillyhatxu/crawler-service/common/logconfig"
	"github.com/sillyhatxu/crawler-service/common/markdown"
	"github.com/sillyhatxu/crawler-service/common/read"
	"github.com/sillyhatxu/crawler-service/common/txt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	project = "crawler-service"
	module  = "algoexpert-go"
)

func init() {
	logconfig.InitialLogConfig(logconfig.Debug(true), logconfig.Env("dev"), logconfig.Project(project), logconfig.Module(module), logconfig.Version("v1.0.0-beta.1"))
}

func algoExpert(filePathName string) bytes.Buffer {
	builder := markdown.CreateBuilder(map[string]markdown.TextType{
		"p":    markdown.Text,
		"h1":   markdown.H1,
		"h2":   markdown.H2,
		"h3":   markdown.H3,
		"h4":   markdown.H4,
		"bold": markdown.Bold,
		"pre":  markdown.CodeBlockQuotes,
	})
	htmlPageByte := read.File(filePathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/h2")
	titleXmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
	title := strings.Trim(strings.ReplaceAll(titleXmlElem.Text, "\n", ""), " ")

	builder.AddContent(markdown.LabelType(markdown.H1), markdown.Content(fmt.Sprintf("%s", title)))

	difficultyXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/div[1]/div[1]/span[2]")
	difficultyXmlElem := colly.NewXMLElementFromHTMLNode(resp, difficultyXmlNode)
	difficulty := difficultyXmlElem.Attr("data-tip")

	builder.AddContent(markdown.Label("bold"), markdown.Content("Difficulty:"))
	builder.AddContent(markdown.Label("difficulty"), markdown.Content(fmt.Sprintf("%s; ", difficulty)))

	categoryXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/div[1]/div[2]/span[2]")
	categoryXmlElem := colly.NewXMLElementFromHTMLNode(resp, categoryXmlNode)
	category := categoryXmlElem.Text

	builder.AddContent(markdown.Label("bold"), markdown.Content("Category:"))
	builder.AddContent(markdown.Label("category"), markdown.Content(fmt.Sprintf("%s\n", category)))

	contentXmlNodes := htmlquery.Find(doc, "/html/body/div/div/div/div[2]/div[2]/div[2]/*")
	for _, contentXmlNode := range contentXmlNodes {
		contentXmlElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
		content := strings.TrimRight(strings.ReplaceAll(contentXmlElem.Text, "  ", ""), "\n")
		if contentXmlElem.Name == "div" {
			break
		}
		builder.AddContent(markdown.Label(contentXmlElem.Name), markdown.Content(fmt.Sprintf("%s", content)))
	}

	hintTitleXmlNode := htmlquery.FindOne(doc, "/html/body/div/div/div/div[2]/div[2]/div[3]/h3")
	hintTitleXmlElem := colly.NewXMLElementFromHTMLNode(resp, hintTitleXmlNode)
	builder.AddContent(markdown.Label(hintTitleXmlElem.Name), markdown.Content(fmt.Sprintf("%s", hintTitleXmlElem.Text)))

	hintXmlNodes := htmlquery.Find(doc, "/html/body/div/div/div/div[2]/div[2]/div[3]/*")
	for _, hintXmlNode := range hintXmlNodes {
		hintXmlElem := colly.NewXMLElementFromHTMLNode(resp, hintXmlNode)
		if hintXmlElem.Name != "div" {
			continue
		}
		hint := hintXmlElem.ChildText("/div/div[1]/div")
		builder.AddContent(markdown.Label("h4"), markdown.Content(fmt.Sprintf("%s", hint)))
		value := strings.TrimRight(strings.ReplaceAll(hintXmlElem.ChildText("/div/div[2]/div/p"), "  ", ""), "\n")
		builder.AddContent(markdown.Label("pre"), markdown.Content(fmt.Sprintf("%s", value)))
	}
	builder.AddContent(markdown.Label(hintTitleXmlElem.Name), markdown.Content(fmt.Sprintf("%s", hintTitleXmlElem.Text)))
	builder.AddContent(markdown.Label("h2"), markdown.Content("Test"))

	testXmlNodes := htmlquery.Find(doc, "/html/body/div/div[3]/div/div[2]/div[3]/*")
	for _, testXmlNode := range testXmlNodes {
		testXmlElem := colly.NewXMLElementFromHTMLNode(resp, testXmlNode)
		testTitle := testXmlElem.ChildText("/div/span")
		testContent := testXmlElem.ChildText("/div/pre")
		builder.AddContent(markdown.LabelType(markdown.H3), markdown.Content(testTitle))
		builder.AddContent(markdown.LabelType(markdown.CodeBlockQuotes), markdown.Content(testContent))
	}
	//fmt.Println("-------------")
	result, err := builder.Build()
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.String())
	return result
}

func main() {
	pwd, _ := os.Getwd()
	for i := 1; i <= 3; i++ {
		fmt.Println(fmt.Sprintf("---------- %d ----------", i))
		fileName := fmt.Sprintf("test%d", i)
		filePathName := fmt.Sprintf("%s/%s/%s.html", pwd, module, fileName)
		logrus.Infof("read file : %s", filePathName)
		result := algoExpert(filePathName)
		f := txt.File{
			FilePath:    fmt.Sprintf("%s/%s/", pwd, module),
			FileName:    fmt.Sprintf("%s.md", fileName),
			FileContent: result,
		}
		err := f.Write()
		if err != nil {
			panic(err)
		}
	}
}
