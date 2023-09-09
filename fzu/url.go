package fzu

import "strconv"

// 新闻页面的URL
func joinURL(i int) string {
	return "https://info22.fzu.edu.cn/content.jsp?urltype=news.NewsContentUrl&wbtreeid=1302&wbnewsid=" +
		strconv.Itoa(i)
}

// 点击数的URL
func joinCountURL(i int) string {
	return "https://info22.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=" +
		strconv.Itoa(i) + "&owner=1768654345&clicktype=wbnews"
}
