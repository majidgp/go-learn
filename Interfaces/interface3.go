package main

import "fmt"

type list1 struct {
	firstname,lastname string
	age,iduser               int
}

func (l *list1)listfull() string {
	return fmt.Sprintf("%s %s %d",l.firstname,l.lastname,l.age)
}
func (al *list1)listid() string {
	return fmt.Sprintf("%d",al.iduser)
}
type list2 struct {
	list1
	idname    int
}

func (li *list2)listfull() string {
	return fmt.Sprintf("%s %s ,id: %d",li.firstname,li.lastname,li.idname)
}

func (bl *list2)listid() string {
	return fmt.Sprintf("%d",bl.idname)
}
type listpk interface {
	listfull() string
}
type listd interface {
	listid() string
}

type dlist interface {
	listpk
	listd
}
func listfu(ls listpk) string {
	return fmt.Sprintf("%s",ls.listfull())
}

func listiu(ld listd) string {
	return fmt.Sprintf("%s",ld.listid())
}

func dblist(db dlist) string {
	return fmt.Sprintf("%s , %s",db.listfull(),db.listid())
}
func main()  {
	listn := &list1{firstname: "majid",lastname: "hp",age: 31,iduser: 1004}
	listm := &list2{}
	listm.firstname = "babak"
	listm.lastname  = "kh"
	listm.idname    = 1001

	//fmt.Println(listfu(listn))
	//fmt.Println(listiu(listn))
	fmt.Println(dblist(listn))
	//fmt.Println(listfu(listm))
	//fmt.Println(listiu(listm))
	fmt.Println(dblist(listm))
}