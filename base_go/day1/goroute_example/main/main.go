package main

import (
	"base_go/day1/goroute_example/goroute"
	"fmt"
)


func main() {
	var pipe chan int
	pipe = make(chan int, 1)

	for i:=0;i <= 5; i++{
		go goroute.Add(i, 300, pipe)
		//sum := <- pipe
		//fmt.Println("sum=", sum)
	}

	for v := range pipe {
		fmt.Println(v)
	}

}