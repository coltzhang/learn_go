package main

import (
	"flag"
	"fmt"
)

type myInt int

type Student struct {
	Name string // 共有成员 可以在包外使用
	age  myInt  // 私有成员 不导出
}

func main() {
	//ar := [...]string{Enone: "no error", Einval: "invalid argument"}	// error
	//println(ar)

	stu := Student{Name: "coltzhang", age: 24}
	fmt.Printf("%v\n", stu)
	println(stu.Name)
	println(stu.age)

	// 获取输入参数 ./jinjie -name zhangchao
	namePtr := flag.String("name", "lyh", "user's name")
	flag.Parse()
	println(*namePtr)
}
