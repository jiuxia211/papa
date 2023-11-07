package main

import (
	"fmt"
	bilibili "papa/bilibli"
	"papa/dal/db"
	"time"
)

func main() {
	db.Init()
	//start := time.Now()
	//fzu.PaFzu()
	//end := time.Now()
	//elapsed := end.Sub(start)
	//fmt.Printf("普通爬取运行时间：%s\n", elapsed)
	//start = time.Now()
	//fzu.PaFzus()
	//end = time.Now()
	//elapsed = end.Sub(start)
	//fmt.Printf("并发爬取运行时间：%s\n", elapsed)
	start := time.Now()
	bilibili.PaBilibiliComment("15072321", "Cookie: buvid3=B35883A6-B110-1D16-B224-C0E0DDE7921938647infoc; b_nut=1693065738; i-wanna-go-back=-1; b_ut=7; _uuid=BED4D323-45B1-936B-E8C2-237C89E9E76D38508infoc; DedeUserID=13528580; DedeUserID__ckMd5=aaab2234010ffb1b; hit-new-style-dyn=1; hit-dyn-v2=1; header_theme_version=CLOSE; LIVE_BUVID=AUTO5516930657715601; rpdid=|(u))kkYu|m~0J'uYmJ|YYumm; buvid4=3B0E9AF3-D9D3-64F2-F71D-7D8EE094FCE012548-022081423-Vk7oLekZ8O%2BXf1iUIja6HA%3D%3D; buvid_fp_plain=undefined; CURRENT_BLACKGAP=0; CURRENT_FNVAL=4048; home_feed_column=5; enable_web_push=DISABLE; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc2NDU0MzYsImlhdCI6MTY5NzM4NjE3NiwicGx0IjotMX0.oLhVAj47786Q4zfCZRnQLE4SyQqrmNm1dDiR8YjzCps; bili_ticket_expires=1697645376; CURRENT_QUALITY=112; SESSDATA=dec2f2e3%2C1713014843%2C3be3f%2Aa2CjDSklemYKgrBp5rYNVlZKDMN7R2vI786BgnbjMK4-ZMBOooHD_EpsARXO8LN-d5iY8SVjJ6aXRMUmNrdDgycGZiNFZWb0xqdTRrWFBMX1hYUmJvbUdxYjc4TTdUTWdvRVB1cDdPcEY0eEdYa3pVZTJoWlA4ZUFxQWQ1Zk00eTdMTDF0SG5kd25RIIEC; bili_jct=14ba10af5817025577be9e44a29d379f; fingerprint=70307773e2bb86a7b65e309ed52954f0; innersign=0; browser_resolution=1536-700; buvid_fp=f9b1a0f81aeab4913ed3a2d48dcd5a75; b_lsid=8316107FA_18B3C84A2BF; bp_video_offset_13528580=853366854503104520; sid=5qdicfbc; PVID=1")
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("普通爬取运行时间：%s\n", elapsed)
}
