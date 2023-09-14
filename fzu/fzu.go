package fzu

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"papa/dal/db"
	"strconv"
	"strings"
	"sync"
)

// 普通爬取
func PaFzu() {
	nextPageURL := "https://info22.fzu.edu.cn/lm_list.jsp?urltype=tree.TreeTempUrl&wbtreeid=1460"

	for {
		resp := fetchHttp(nextPageURL)
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// 获取该页中所有的通知URL
		newsURL := getNewsURL(doc)
		// 爬取每个通知的详细信息
		for _, URL := range newsURL {
			resp := fetchHttp(URL)
			parse(resp)
		}
		nextPageURL = "https://info22.fzu.edu.cn/lm_list.jsp?" + getNextPageURL(doc)
		//log.Printf("下一页的URL为%v", nextPageURL)
		parsedURL, err := url.Parse(nextPageURL)
		if err != nil {
			fmt.Println("解析URL时发生错误:", err)
			return
		}
		// 爬到50页停止
		pageNum, _ := strconv.Atoi(parsedURL.Query().Get("PAGENUM"))
		if (pageNum) == 50 {
			log.Println("已经爬到最后一页了")
			break
		}
	}

}

// 并发爬取
func PaFzus() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // 控制并发数为 5
	nextPageURL := "https://info22.fzu.edu.cn/lm_list.jsp?urltype=tree.TreeTempUrl&wbtreeid=1460"
	for {
		resp := fetchHttp(nextPageURL)
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// 获取该页中所有的通知URL
		newsURL := getNewsURL(doc)
		// 爬取每个通知的详细信息
		for _, URL := range newsURL {
			wg.Add(1)
			// 创建中间变量
			URL := URL
			go func() {
				sem <- struct{}{}        // 获取一个信号量，限制并发数
				defer func() { <-sem }() //  // 释放一个信号量
				defer func() { wg.Done() }()
				resp := fetchHttp(URL)
				parse(resp)
			}()

		}
		wg.Wait()
		nextPageURL = "https://info22.fzu.edu.cn/lm_list.jsp?" + getNextPageURL(doc)
		//log.Printf("下一页的URL为%v", nextPageURL)
		parsedURL, err := url.Parse(nextPageURL)
		if err != nil {
			fmt.Println("解析URL时发生错误:", err)
			return
		}
		// 爬到50页停止
		pageNum, _ := strconv.Atoi(parsedURL.Query().Get("PAGENUM"))
		if (pageNum) == 50 {
			log.Println("已经爬到最后一页了")
			break
		}
	}
}
func getNewsURL(doc *goquery.Document) []string {
	newsURLs := make([]string, 0)
	newsURLSelector := "body > div.sy-content > div.content > div.right.fr > div.list.fl > ul > li"
	liElements := doc.Find(newsURLSelector)

	liElements.Each(func(i int, liElement *goquery.Selection) {
		aElements := liElement.Find("p a")
		if aElements.Length() >= 2 {
			newsURL, exist := aElements.Eq(1).Attr("href")
			if !exist {
				log.Printf("未找到通知文件目录的通知URL")
			}
			newsURLs = append(newsURLs, "https://info22.fzu.edu.cn/"+newsURL)
		}

	})
	return newsURLs
}
func getNextPageURL(doc *goquery.Document) string {
	nextPageElements := doc.Find("body > div.sy-content > div > div.right.fr " +
		"> div.list.fl > div > span.p_pages > span.p_next.p_fun > a")
	nextPageURL, exist := nextPageElements.Attr("href")
	if !exist {
		log.Println("获取下一页URL失败")
		return ""
	} else {
		return nextPageURL
	}
}
func fetchHttp(URL string) (resp *http.Response) {
	client := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
func parse(resp *http.Response) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var titleBuilder strings.Builder
	// 标题的css选择器路径
	titleSelector := "body > div.wa1200w > div.conth > form > div.conth1"
	titleElement := doc.Find(titleSelector)
	titleBuilder.WriteString(titleElement.Text())
	var dataBuilder strings.Builder
	// 日期的css选择器路径
	dateSelector := "body > div.wa1200w > div.conth > form > div.conthsj"
	dateElement := doc.Find(dateSelector)
	parts := strings.Split(dateElement.Text(), "：")
	if len(parts) > 1 {
		dataBuilder.WriteString(parts[1])
	}
	var contentBuilder strings.Builder
	// 内容的css选择器路径
	contentSelector := "#vsb_content div.v_news_content p"
	pElements := doc.Find(contentSelector)
	pElements.Each(func(i int, pElement *goquery.Selection) {
		// 在当前 <p> 元素下查找所有的 <span> 元素
		spanElements := pElement.Find("span")
		spanElements.Each(func(i int, spanElement *goquery.Selection) {
			contentBuilder.WriteString(spanElement.Text())
		})
	})
	//剔除一些空页面
	if titleBuilder.String() == "" || contentBuilder.String() == "" || dataBuilder.String() == "" {
		return
	} else {
		err := db.CreateNews(&db.News{
			Title:   titleBuilder.String(),
			Content: contentBuilder.String(),
			Date:    dataBuilder.String(),
		})
		if err != nil {
			log.Print("创建news失败")
		}
	}
}
