package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	//sli := make([]string, 3)
	//
	//fmt.Println(len(sli), cap(sli))
	//sli[0] = "张三"
	//sli[1] = "李四"
	//sli[2] = "王五"

	sli := make([]string, 0)

	fmt.Println(len(sli), cap(sli))

	sli = append(sli, "张三")
	sli = append(sli, "李四")
	sli = append(sli, "王五")

	//var sli = []string{"张三", "李四", "王五"}

	fmt.Println(len(sli), cap(sli))

	fmt.Printf("main中sli的内存地址%p:", sli)
	spew.Dump("main中sli的内存地址:", sli)

	f11(sli)

	fmt.Println(sli)

}

func f11(arg []string) {

	fmt.Printf("f1中参数的内存地址%p:", arg)
	spew.Dump("f1中参数的内存地址:", arg)

	arg[1] = "彭齐"
	//arg[3] = "彭齐"
	//arg[4] = "彭齐"
	arg = append(arg, "爽哥1")
	arg = append(arg, "爽哥2")
	arg = append(arg, "爽哥3")

}
