package main

import (
	"fmt"
	"papa/dal/db"
	"papa/fzu"
	"time"
)

func main() {
	db.Init()
	start := time.Now()
	fzu.PaFzu()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("普通爬取运行时间：%s\n", elapsed)
	start = time.Now()
	fzu.PaFzus()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("并发爬取运行时间：%s\n", elapsed)
}
