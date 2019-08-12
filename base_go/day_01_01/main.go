package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// 并发版

func HttpGet(url string) (result string, err error) {

	resp , err1 := http.Get(url)
	if err1 != nil{
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024*4)

	for {
		n ,err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.body err", err)
			break
		}
		result+=string(buf[:n])
	}
	return

}


////爬取一个网页
//func SpiderPape(i int, page chan<- int) {
//	url := "https://tieba.baidu.com/p/6213420290?pn=" + strconv.Itoa(i)
//	fmt.Printf("正在爬第%d页网页: %s\n", i, url)
//
//	//2) 爬 (将所有的网站的内容全部爬下来)
//	result, err := HttpGet(url)
//	if err != nil {
//		fmt.Println("HttpGet err = ", err)
//		return
//	}
//
//	//把内容写入到文件
//	fileName := strconv.Itoa(i) + ".html"
//	f, err1 := os.Create(fileName)
//	if err1 != nil {
//		fmt.Println("os.Create err1 = ", err1)
//		return
//	}
//
//	f.WriteString(result) //写内容
//
//	f.Close() //关闭文件
//
//	page <- i
//}
//
//func DoWork(start, end int) {
//	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)
//	page := make(chan int)
//
//	//明确目标 (要知道你准备在哪个范围或者网站去搜索)
//	//http://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0  //下一页+50
//	for i := start; i <= end; i++ {
//		go SpiderPape(i, page)
//	}
//
//	for i := start; i <= end; i++ {
//		fmt.Printf("第%d个页面爬取完成\n", <-page)
//	}
//
//}


func SpiderWrite(i int, page chan<- int)  {
	//生成url
	url := "https://tieba.baidu.com/p/6213420290?pn="+ (strconv.Itoa(i))
	fmt.Printf("%s\n", url)
	// scripy
	result, err := HttpGet(url)
	if err != nil{
		fmt.Println("http err", err)
	}
	// write
	fileName := strconv.Itoa(i)+".html"
	f, err1 := os.Create(fileName)
	if err1 != nil{
		fmt.Println("write err", err1)
	}
	f.WriteString(result)
	f.Close()
	page <- i

}

func DoWork(start, end int)  {

	fmt.Printf("正在爬取%d 到 %d\n", start, end)

	page := make(chan int)

	for i:= start; i<=end; i++{
		go SpiderWrite(i, page)
	}

	for i:= start; i<=end; i++{
		fmt.Printf("第%d个", <- page)}

}


func main(){

	var start, end int

	fmt.Println("qishi:")
	fmt.Scan(&start)
	fmt.Println("end")
	fmt.Scan(&end)

	DoWork(start, end)

}