package main

import "fmt"

func main() {

	s := MyStruct{x: 0}
	fmt.Printf("初始化后的地址----------> %p\n", &s.x)
	s.Set1()
	fmt.Println(s.x)
	s.Set2()
	fmt.Println(s.x)
}
