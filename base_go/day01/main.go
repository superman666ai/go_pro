package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

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

func DoWork(start, end int)  {

	fmt.Printf("正在爬取%d 到 %d\n", start, end)

	for i:= start; i<=end; i++{
		//生成url
		url := "https://tieba.baidu.com/p/6213420290?pn="+ (strconv.Itoa(i))
		fmt.Printf("%s\n", url)
		// scripy
		result, err := HttpGet(url)
		if err != nil{
			fmt.Println("http err", err)
			continue
		}
		// write
		fileName := strconv.Itoa(i)+".html"
		f, err1 := os.Create(fileName)
		if err1 != nil{
			fmt.Println("write err", err1)
			continue
		}
		f.WriteString(result)
		f.Close()
	}
	
}
func main(){

	var start, end int

	fmt.Println("qishi:")
	fmt.Scan(&start)
	fmt.Println("end")
	fmt.Scan(&end)

	DoWork(start, end)

}