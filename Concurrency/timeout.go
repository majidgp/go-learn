package main

import (
	"fmt"
	"time"
)

func main()  {

	channel := make(chan int)

	go func() {
		for  {
			channel <- 100
			time.Sleep(1999 * time.Millisecond)
		}
	}()
	for  {
		select {
		case n := <- channel:
			fmt.Println(n)
		case <- time.After(2000 * time.Millisecond):
			fmt.Println("Timed out!!!")
			return
		}

	}
}
