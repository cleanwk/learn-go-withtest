package main

import "fmt"

type MyStruct struct {
	x int
}

func (m MyStruct) Set1() {
	fmt.Printf("recevier值传递x的地址---> %p\n", &m.x)
	m.x = 1
}

func (m *MyStruct) Set2() {
	fmt.Printf("recevier引用传递x地址---> %p\n", &m.x)
	m.x = 2
	fmt.Println(&m.x)

}
