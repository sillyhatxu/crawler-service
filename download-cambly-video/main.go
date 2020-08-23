package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/sillyhatxu/crawler-service/common/logconfig"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	project = "crawler-service"
	module  = "download-cambly-video"
)

func init() {
	logconfig.InitialLogConfig(logconfig.Debug(true), logconfig.Env("dev"), logconfig.Project(project), logconfig.Module(module), logconfig.Version("v1.0.0-beta.1"))
}

type Cambly struct {
	Tutor string
	Date  string
	URL   string
}

func main() {
	pwd, _ := os.Getwd()
	fileName := "test.html"
	filePathName := fmt.Sprintf("%s/%s/%s", pwd, module, fileName)
	cookiePathName := fmt.Sprintf("%s/%s/%s", pwd, module, "cookie.txt")
	hostPathName := fmt.Sprintf("%s/%s/%s", pwd, module, "host.txt")
	logrus.Infof("read file : %s", filePathName)
	htmlPageByte := readFile(filePathName)
	cookie := readFile(cookiePathName)
	host := readFile(hostPathName)
	resp := &colly.Response{StatusCode: 200, Body: htmlPageByte}
	doc, _ := htmlquery.Parse(strings.NewReader(string(htmlPageByte)))
	xmlNodes := htmlquery.Find(doc, "/html/body/div/div/div")
	var camblyArray []Cambly
	for _, xmlNode := range xmlNodes {
		xmlElem := colly.NewXMLElementFromHTMLNode(resp, xmlNode)
		//fmt.Print("name : ", xmlElem.ChildText("/div/div/div/h3/a"))
		//fmt.Print("; date : ", formatTime(xmlElem.ChildText("/div/div/div/p")))
		//fmt.Print("; url : ", xmlElem.ChildAttr("/div[2]/div/ul/li[2]/a", "href"))
		//fmt.Println()
		url := xmlElem.ChildAttr("/div[2]/div/ul/li[2]/a", "href")
		if url != "" {
			url = string(host) + xmlElem.ChildAttr("/div[2]/div/ul/li[2]/a", "href")
		}
		camblyArray = append(camblyArray, Cambly{
			Tutor: xmlElem.ChildText("/div/div/div/h3/a"),
			Date:  formatTime(xmlElem.ChildText("/div/div/div/p")),
			URL:   url,
		})
	}
	fmt.Println(len(camblyArray))
	for i, cambly := range camblyArray {
		url := redirect(cambly.URL, string(cookie))
		fileName := getFileName(cambly)
		fmt.Println("download file[", i, "] start : ", fileName, url)
		if cambly.URL == "" {
			fmt.Println("download fault : ", fileName, cambly.URL)
			continue
		}
		err := downloadFile(fmt.Sprintf("/Users/shikuanxu/Documents/Video/Cambly/%s", fileName), url)
		if err != nil {
			fmt.Println("download fault : ", fileName, url)
			continue
		}
		fmt.Println("download success : ", fileName, cambly.URL)
		time.Sleep(10 * time.Second)
	}
}

var check = make(map[string]int)

func getFileName(cambly Cambly) string {
	name := fmt.Sprintf("%s(%s)", cambly.Date, cambly.Tutor)
	result, ok := check[name]
	if ok {
		check[name] = result + 1
		return fmt.Sprintf("%s(%d)", name, check[name]) + ".mp4"
	}
	check[name] = 1
	return name + ".mp4"
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	contextByte, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return contextByte
}

func formatTime(input string) string {
	format, err := time.Parse("Jan 2, 2006 3:04 PM", input)
	if err != nil {
		return ""
	}
	return format.Format("2006-01-02")
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func redirect(url string, cookie string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Cookie", cookie)
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	return res.Request.URL.String()
	//client := &http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		return http.ErrUseLastResponse
	//	},
	//}
	//c := http.Cookie{
	//	Name: "theme",
	//	Value: "dark",
	//}
	//http.SetCookie(w, &c)
	//res, err := client.Get(url)
	//if err != nil {
	//	return ""
	//}
	//if res.StatusCode != 301 {
	//	return ""
	//}
	//return res.Header.Get("Location")
}
