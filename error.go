package main

import (
	"fmt"
	"strconv"
	"time"
)

type errorCommon struct {
	ErrCode int
	ErrMsg  string
}

func (e *errorCommon) Error() map[string]string {
	return map[string]string{
		"ErrCode": strconv.Itoa(e.ErrCode),
		"ErrMsg":  e.ErrMsg,
	}
}

func main() {
	//i, err := strconv.Atoi("a")

	//var err interface{} = &errorCommon{
	//	ErrCode: 404,
	//	ErrMsg:  "无法找到",
	//}
	//if com, ok := err.(*errorCommon); ok {
	//	fmt.Println(com.ErrCode, com.ErrMsg)
	//	fmt.Println(err == os.ErrExist)
	//	w := fmt.Errorf("错误哦%w", com)
	//	fmt.Println(w)
	//	fmt.Println(errors.Unwrap(w), err == os.ErrExist, errors.Is(w, os.ErrExist))
	//} else {
	//	fmt.Println(ok, com)
	//}

	go errorsdef()
	go fmt.Println("1")
	go fmt.Println("2")
	fmt.Println("3")
	time.Sleep(1 * time.Second)
}

func errorsdef() {
	defer fmt.Println("4")
	defer fmt.Println("5")
	defer fmt.Println("6")
}
