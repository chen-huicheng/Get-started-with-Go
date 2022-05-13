package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Personer interface {
	Humaner
	sing(lrc string)
}
type Studert struct {
	name string
	id   int
}

func (this *Studert) sayhi() {
	fmt.Printf("student %+v\n", *this)
}

type Teacher struct {
	name string
	id   int
}

func (this *Teacher) sayhi() {
	fmt.Printf("teacher %+v\n", this)
}

func WhoSayHi(i Humaner) {
	i.sayhi()
}
func main() {
	var i Humaner
	//只要实现了该接口方法的类型，那么这个类型的变量就可以该i赋值

	//cannot use s2 (variable of type Studert) as Humaner value in assignment: missing method sayhi (sayhi has pointer receiver)
	// s2 := Studert{"w", 5}  接口 i sayhi方法的receiver是 *Student 所以必须使用 i = &s2
	// i = s2

	i = &Studert{"c", 4}
	i.sayhi()

	s := Studert{"chen", 24}
	t := Teacher{"me", 24}
	WhoSayHi(&s)
	WhoSayHi(&t)
}
