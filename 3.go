package main

import "fmt"

func main() {
	a := []byte{1, 0}
	fmt.Println("len of old a is ", len(a)) // 2
	fmt.Println("cap of old a is ", cap(a)) // 2
	fmt.Println("")

	a = append(a, 1, 1, 1)
	fmt.Println("len of a is ", len(a)) // 5
	fmt.Println("cap of a is ", cap(a)) // 8

	fmt.Println("------")

	b := []int{23, 51}
	fmt.Println("len of old b is ", len(b)) // 2
	fmt.Println("cap of old b is ", cap(b)) // 2
	fmt.Println("")

	b = append(b, 4, 5, 6)
	fmt.Println("len of b is ", len(b)) // 5
	fmt.Println("cap of b is ", cap(b)) // 6

	fmt.Println("------")

	c := []int32{1, 23}
	fmt.Println("len of old c is ", len(c)) // 2
	fmt.Println("cap of old c is ", cap(c)) // 2
	fmt.Println("")

	c = append(c, 2, 5, 6)
	fmt.Println("len of c is ", len(c)) // 5
	fmt.Println("cap of c is ", cap(c)) // 6

	fmt.Println("------")

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}
	fmt.Println("len of old d is ", len(d)) // 2
	fmt.Println("cap of old d is ", cap(d)) // 2
	fmt.Println("")

	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"})
	fmt.Println("len of d is ", len(d)) // 5
	fmt.Println("cap of d is ", cap(d)) // 5

}
