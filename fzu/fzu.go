package fzu

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"papa/dal/db"
	"strings"
	"sync"
)

// 普通爬取
func PaFzu() {
	for i := 33638; i >= 33008; i-- {
		URL := joinURL(i)
		resp := fetchHttp(URL)

		parse(resp)
	}
}

// 并发爬取
func PaFzus() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // 控制并发数为 5
	for i := 33638; i >= 33008; i-- {
		wg.Add(1)
		index := i // 局部副本的变量,避免使用错误的i值
		go func() {
			sem <- struct{}{}        // 获取一个信号量，限制并发数
			defer func() { <-sem }() //  // 释放一个信号量
			URL := joinURL(index)
			resp := fetchHttp(URL)
			parse(resp)
		}()
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
