package main

import "fmt"

type FuncType func(int, int) int

type Student struct {
	id     int
	name   string
	gender byte
	age    int
	addr   string
}

func main() {
	var s1 Student = Student{1, "chen", 0, 25, "sy"}
	fmt.Println(s1)

	s2 := Student{id: 2, name: "wang"}
	fmt.Println(s2)

	var p *Student
	p = &s1
	fmt.Println(*p)

	//成员使用
	s := Student{}
	s.id = 3
	s.name = "me"
	fmt.Println(s.name)

	//指针

	p.id = 3
	(*p).name = "cheng"
	fmt.Println(*p)

	p2 := new(Student)
	p2.id = 4
	p2.name = "he"

	//比较 和赋值
	s3 := s1
	fmt.Println(s3)
	fmt.Println(s3 == s1)
	fmt.Println(s2 == s1)

	s2 = s3
	fmt.Println(s2 == s3)
}
