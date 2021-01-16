package main

import (
	"fmt"
	"time"
)

func mytext(text string)  {
	for i:=0;i<5;i++{
		fmt.Println(text)
		time.Sleep(500 * time.Millisecond)
	}
}

func main()  {
	go mytext("hello")
	mytext("majid")
}