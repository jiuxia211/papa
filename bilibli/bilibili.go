package bilibili

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"papa/dal/db"
	"time"
)

func PaBilibiliComment(oid string, cookie string) {
	pn := 1
	// url := "https://api.bilibili.com/x/v2/reply?type=1&oid=21071819&sort=0&pn=1"
	for {
		Url := fmt.Sprintf("https://api.bilibili.com/x/v2/reply?type=1&oid=%v&sort=0&pn=%v", oid, pn)
		resp := fetchHttpWithCookie(Url, cookie)
		if end := parseParent(resp, oid, cookie); end {
			break
		}
		pn++
	}
}
func fetchHttpWithCookie(URL string, cookie string) (resp *http.Response) {
	client := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("Cookie", cookie)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// 解析返回并存入数据库
func parseParent(resp *http.Response, oid string, cookie string) (end bool) {
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("将resp转化为string时发生错误:", err)
		return
	}
	replies := gjson.Get(string(bodyText), "data.replies")
	replyArray := replies.Array()
	if len(replyArray) == 0 {
		return true
	}
	for _, reply := range replyArray {
		content := reply.Get("content.message").String()
		like := reply.Get("like").Int()
		replyTime := reply.Get("ctime").Int()
		rpid := reply.Get("rpid").Int()
		rcount := reply.Get("rcount").Int()
		err := db.CreateParentReply(&db.ParentReply{
			Content:   content,
			Like:      like,
			ReplyTime: replyTime,
			Rpid:      rpid,
		})
		if err != nil {
			log.Print("创建reply失败")
		}
		if rcount != 0 {
			pn := 1
			for {
				Url := fmt.Sprintf("https://api.bilibili.com/x/v2/reply/reply?type=1&oid=%v&root=%v&pn=%v", oid, rpid, pn)
				childResp := fetchHttpWithCookie(Url, cookie)
				if childEnd := parseChild(childResp); childEnd {
					break
				}
				pn++
				time.Sleep(time.Second)
			}

		}
	}
	return false
}

// 解析子评论并存入数据库
func parseChild(resp *http.Response) (end bool) {
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("将resp转化为string时发生错误:", err)
		return
	}
	replies := gjson.Get(string(bodyText), "data.replies")
	replyArray := replies.Array()
	if len(replyArray) == 0 {
		return true
	}
	for _, reply := range replyArray {
		content := reply.Get("content.message").String()
		like := reply.Get("like").Int()
		replyTime := reply.Get("ctime").Int()
		rpid := reply.Get("rpid").Int()
		parentID := reply.Get("root").Int()
		err := db.CreateChildReply(&db.ChildReply{
			Content:   content,
			Like:      like,
			ReplyTime: replyTime,
			Rpid:      rpid,
			ParentID:  parentID,
		})
		if err != nil {
			log.Print("创建reply失败")
		}
	}
	return false
}
