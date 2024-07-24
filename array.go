package main

import "fmt"

func main() {
	//a := [3]int{1, 2, 3}
	//b := [4]int{1, 2, 3}
	//c := []int{1, 2, 3}
	//d := []int{1, 2, 3, 4, 5}
	//d = c
	//fmt.Printf("a: %T\n", a)
	//fmt.Printf("b: %T\n", b)
	//fmt.Printf("c: %T\n", c)
	//fmt.Printf("d: %T\n", d)
	//
	//for i, v := range d {
	//	fmt.Println(i, v)
	//}

	//e := map[int]map[int]string{
	//	1: {
	//		1: "a",
	//		2: "b",
	//	},
	//	2: {
	//		1: "c",
	//		2: "d",
	//	},
	//}

	//f := func(a map[int]map[int]string) int {
	//	sum := 0
	//	for i, _ := range a {
	//		sum += i
	//	}
	//	return sum
	//}
	//fmt.Println(f(e))

	//s := stu(1)
	//fmt.Println(s, *&s, &s)
	//s.string()
	//s = stu(234)
	//b := &s
	//*b = stu(345)
	//fmt.Println(s, &s)

	//stu.strings(stu(123))
	t := stu(10)
	t.string()
	s := stu.strings
	s(stu(10))
}

type stu int
type sdu stu

func (s *stu) string() {
	*s = stu(123)
	*s = stu(sdu(780))
	fmt.Println(*s, &s)
}
func (s stu) strings() {
	fmt.Println(s)
}
