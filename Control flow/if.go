package main

import (
	"errors"
	"fmt"
)

func foo()(err error)  {
	return errors.New("error")
}
func main()  {

	number := 20
	if number >18{
		fmt.Println("your welcome to website")
	}else if number <18{
		fmt.Println("sorry")
	}else {
		fmt.Println("limited access to website")
	}
	if err := foo();err != nil{
		fmt.Println("we have on error here")

	}

}
