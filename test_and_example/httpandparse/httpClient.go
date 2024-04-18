package httpandparse

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/imroc/req/v3"
	"github.com/liuzl/gocc"
	"golang.org/x/net/html"
)

type LoadConfig struct {
}

func Load() {
	TestHttpClient("https://www.ibiquges.org/92/92872/")
}

func TestHttpClient(url string) {

	t2s, err := gocc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}

	client := req.C() // Use C() to create a client.
	request := client.R()

	//header
	headerMap := map[string]string{
		"Accept":     " application/json, text/javascript, */*; q=0.01",
		"Referer":    "https://czbooks.net/",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}
	request.SetHeaders(headerMap)

	// Get
	resp, err := request.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(resp)

	// 解析为Node
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("html.Parse(resp.Body) err: %+v", err)
	}
	//xpath选择
	//list, err := htmlquery.QueryAll(doc, "//*[@id=\"chapter-list\"]/*")
	li, _ := htmlquery.QueryAll(doc, "//*[@id=\"chapter-list\"]/li")
	for i := 0; i < len(li); i++ {
		element := li[i]
		if element.Data == "li" && element.FirstChild.Data == "a" {
			attrs := element.FirstChild.Attr
			for i := 0; i < len(attrs); i++ {
				attr := attrs[i]
				if attr.Key == "href" {
					pageUrl := "https:" + attr.Val
					r, _ := request.Get(pageUrl)
					// 解析为Node
					rdoc, err := goquery.NewDocumentFromReader(r.Body)
					if err != nil {
						log.Fatal(err)
					}
					//selector 的语法
					rnodes := rdoc.Find("#sticky-parent > div.chapter-detail.style-left > div.content").Text()
					simpleContent, err := t2s.Convert(rnodes)
					if err != nil {
						fmt.Printf("t2s.Convert(resp.Body) err: %+v", err)
					}
					c := fmt.Sprintf("\n%s\n", simpleContent)
					writeToFile(c)
				}
			}
		}
	}
}

func writeToFile(c string) {
	fileN := "农家子的古代科举生活.txt"
	filePrefix, _ := filepath.Abs(".")
	fileP := filePrefix + "/" + fileN
	// 可写方式打开文件
	file, err := os.OpenFile(
		fileP,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		fs.ModePerm,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	bytesWritten, err := file.Write([]byte(c))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote %+V bytes.\n", bytesWritten)
}
