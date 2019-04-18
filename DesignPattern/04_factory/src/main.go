package main

import "fmt"

type interA interface {
	f1()
	f2()
}

type structA struct {
}

func (a *structA) f1() {
	fmt.Println("this is f1.default")
}

func (a *structA) f2() {
	fmt.Println("this is f2")
	a.f1()
}

type structB struct {
	structA
}

func (a *structB) f1() {
	fmt.Println("this is f1.v2")
}

func main() {
	var b structB
	b.f2()
}
