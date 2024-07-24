package main

import "fmt"

type stringer interface {
	Strings(string2 string) string
}

type stua struct {
	id   int
	name string
}

type address struct {
	city    string
	country string
}

func main() {
	a := stua{id: 1, name: "小子"}
	//b := address{
	//	city:    "杭州",
	//	country: "中国",
	//}
	//fmt.Println(a.String())

	printString(&a)
	var b interface{} = &a
	//c := (*b).(*stua)
	c, ok := b.(*stua)
	fmt.Printf("b: %T\n", c, ok, b)
	//printString(b)
	//struccst(a)
}

func (s *stua) Strings(str string) string {
	s.name += str
	return s.name
}

func (a address) Strings(str string) string {
	return a.country + a.city + str
}

func printString(s stringer) {
	fmt.Println(s.Strings("waoun"))
}

func struccst(a stua) {
	fmt.Println(a.Strings("waoun"))
}
