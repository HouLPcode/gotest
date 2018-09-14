package main

//main函数没有结束，子函数中创建的goroutine不会被强制回收

import (
	"fmt"
	"time"
)

func main(){
	func1()
	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}

func func1(){
	go func() {
		fmt.Println("i am a new goroutine")
		time.Sleep(time.Second*5)
		fmt.Println("goroutine over")
	}()
	defer func() {
		fmt.Println("defer func1")
	}()
}
