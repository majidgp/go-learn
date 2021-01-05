package main

import "fmt"

type artist struct {
	name,genre string
	songs      int
}

func newrelease(a *artist)int  {
     a.songs++
     return a.songs
}

func main()  {
	me := &artist{"majid","electro",42}
	fmt.Println(me.name,me.genre,newrelease(me))
	fmt.Println(me.name,me.genre,me.songs)
}
