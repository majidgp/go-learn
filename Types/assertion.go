package main

import (
	"fmt"
	"time"
)

func timemap(x interface{})  {

	y,ok := x.(map[string]interface{})
	if ok{
		y["Update_now"] = time.Now()
	}

}
func main()  {

	timed := map[string]interface{}{
		"majid":30,
	}
	timemap(timed)
	fmt.Println(timed)

}