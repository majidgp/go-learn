package main

import (
	"fmt"
	"time"
)

func sendvl(channel chan int,vl int)  {
	for i:=0;i<10;i++{
		time.Sleep(time.Millisecond * 1000)
		channel <- vl *i
	}
}
func recvl(channel chan int)  {
	for{
		vl := <- channel
		fmt.Println(vl)
	}
}
func main()  {
	channel1 := make(chan int,20)
	vl1 := 1000
	go sendvl(channel1,vl1)
	go recvl(channel1)
	for  {
		
	}
}
