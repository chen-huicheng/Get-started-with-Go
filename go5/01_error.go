package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (ret int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		ret = a / b
	}
	return
}

func main() {
	err1 := fmt.Errorf("%s", "this is normol err")

	fmt.Println(err1)

	err2 := errors.New("hello error")
	fmt.Println(err2)

	//应用
	res, err := MyDiv(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
