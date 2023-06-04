package main

import "fmt"

func main() {

	sli := []int{1, 1, 2, 3, 4}
	sl2 := []int{2, 3, 4, 1}

	rs := DiffArray(sli, sl2)

	fmt.Println(rs)
}

// DiffArray 求两个切片的差集
func DiffArray(a []int, b []int) []int {
	var diffArray []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}
