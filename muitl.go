package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("len(s1):%d,cap(s1):%d\n", len(s1), cap(s1))
	s1 = append(s1, "广州", "成都", "重庆", "石家庄", "保定", "邢台", "张家口", "济南")
	fmt.Printf("len(s1):%d,cap(s1):%d\n", len(s1), cap(s1))
}
