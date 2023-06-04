package main

import (
	"fmt"
)

func main() {
	sli := []float64{1, 2, 3.14}

	rs := InSlice(3.14, sli)
	fmt.Println(rs)

}

func InSlice[T int | int8 | int32 | int64 | float32 | float64 | string](ele T, sli []T) bool {
	for _, v := range sli {
		if v == ele {
			return true
		}
	}
	return false
}
