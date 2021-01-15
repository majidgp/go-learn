package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	createdat   time.Time
	body        string
}

func (me *Myerror)Error() string {
    return fmt.Sprintf("At %s Error: %s",me.createdat,me.body)
}

func RunApp() error {
	newerror := &Myerror{createdat: time.Now(),body: "we have funny error here"}
	return newerror
}
func main()  {
	myapp := RunApp()
	fmt.Println(myapp)
}