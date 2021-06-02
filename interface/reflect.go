package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a int "tag"
	b int
}

func main() {
	//var n int = 3
	n := S{1, 2}
	t := reflect.TypeOf(n)
	v := reflect.ValueOf(n)
	fmt.Println(t) // 自定义的类型
	fmt.Println(t.Name())
	fmt.Println(t.Kind()) // 内置类型
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())

	if v.Kind() == reflect.Struct {
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("type:%T, value:%v\n", v.Field(i), v.Field(i))
		}
	}

	//fmt.Println(v.Elem().Field(0))	// 指针才行
	//fmt.Println(v.Elem().Field(1))

}
