package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/sillyhatxu/crawler-service/common/read"
	"os"
	"strings"
)

const (
	project = "crawler-service"
	module  = "dy10000"
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
	//resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	titleXmlNodes := htmlquery.Find(doc, "/html/body/ul/*/span[2]/input")
	for _, titleXmlNode := range titleXmlNodes {
		//xmlElem := colly.NewXMLElementFromHTMLNode(resp, titleXmlNode)
		//contentXmlNode := htmlquery.FindOne(titleXmlNode, "/span[2]/input")
		//fmt.Println(contentXmlNode.Data)
		for _, attribute := range titleXmlNode.Attr {
			if attribute.Key == "value" {
				fmt.Println(attribute.Val)
			}
		}

		//	title := replace(xmlElem.ChildText("/h2"))
		//	titleContent := strings.Split(title, "-")
		//	questionCount := strings.Split(titleContent[1], "/")
		//	//fmt.Println(fmt.Sprintf("%d. %s-%s(%s)", i+1, titleContent[0], questionCount[1], "Easy"))
		//	folder1 := fmt.Sprintf("%d. %s(%s)", i+1, titleContent[0], questionCount[1])
		//	fmt.Println(folder1)
		//	_ = os.Mkdir("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/test/"+folder1, 0755)
		//	contentXmlNodes := htmlquery.Find(titleXmlNode, "/div/*")
		//	index := 1
		//	for j, contentXmlNode := range contentXmlNodes {
		//		contentElem := colly.NewXMLElementFromHTMLNode(resp, contentXmlNode)
		//		difficulty := "Unknown"
		//		if j%2 == 1 {
		//			diffXmlNode := htmlquery.FindOne(contentXmlNode, "/span")
		//			diffElem := colly.NewXMLElementFromHTMLNode(resp, diffXmlNode)
		//			if diffElem.Attr("class") == "_1pp4EraVnLQtN6OfzsaR4j Gy9P8MyDDSdmBqEs4ShcI" {
		//				difficulty = "Easy"
		//			} else if diffElem.Attr("class") == "_1pp4EraVnLQtN6OfzsaR4j STrLX2dE7KQLix7FlqXPy" {
		//				difficulty = "Medium"
		//			} else if diffElem.Attr("class") == "_1pp4EraVnLQtN6OfzsaR4j _1fKglLjhMc0eS6_rvbvA49" {
		//				difficulty = "Hard"
		//			} else if diffElem.Attr("class") == "_1pp4EraVnLQtN6OfzsaR4j Q9gJKSpVldmU7eIN1Lp9H" {
		//				difficulty = "Very Hard"
		//			}
		//			subTitle := replace(contentElem.Text)
		//			folder2 := fmt.Sprintf("%d.%d %s(%s)", i+1, index, subTitle, difficulty)
		//			fmt.Println(folder2)
		//			_ = os.Mkdir("/Users/shikuanxu/go/src/github.com/sillyhatxu/crawler-service/test/"+folder2, 0755)
		//			index++
		//		}
		//	}
	}
}
