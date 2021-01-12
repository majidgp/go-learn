package main

import "fmt"

type Newstruct struct {
	name   string
	family string
	model  newmodel
}

type newmodel struct {
	id      uint64
	address string
}

func (s *Newstruct)Sayhello() string {
      return  "hello"
}

func main()  {
	newstruct := Newstruct{
		name : "majid",
		family : "hp",
		model: newmodel{
			id : 21,
			address: "shahi",
		},
	}
	newstruct.Sayhello()
	fmt.Println(newstruct.name)
	fmt.Println(newstruct.model.id)
}