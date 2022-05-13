package main

import "fmt"

func test1(x int) {

	defer func() {
		//recover() //当发生 panic 错误时 恢复错误
		if err := recover(); err != nil {
			fmt.Println(err)
		}

	}()
	var a [10]int
	a[x] = 11
	fmt.Println("set successful")

}
func main() {
	// test()
	test1(51)
	//panic
	//"runtime error: index out of range [15] with length 10"

}
