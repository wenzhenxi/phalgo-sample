package main

import (
	"time"
	"fmt"
)

var chs chan int = make(chan int,1)

func write(){
	time.Sleep(3*time.Second)
	chs<-1
}

func read(){
	select {
	case ch1:=<-chs:
		fmt.Println(ch1)
		return
	case <-time.After(4*time.Second):
		fmt.Println("read time out")
		return
	}
}

func main() {
	go write()
	read()
}