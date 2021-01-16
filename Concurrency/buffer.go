package main

import "fmt"

func main()  {
	cl := make(chan int,2)
	cl <- 1
	cl <- 2
	cl3 := func() {cl <- 3}
	go cl3()
	fmt.Println(<-cl)
	fmt.Println(<-cl)
	fmt.Println(<-cl)
}
