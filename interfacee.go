package main

import (
	"fmt"
)

func main() {
	sli := []int64{1, 2, 3}
	var sliIface []interface{}

	for _, item := range sli {
		sliIface = append(sliIface, item)
	}

	rs := InSliceIface(int64(2), sliIface) // 2 必须指定为int64类型，否则会当成int，最终结果为false
	fmt.Println(rs)
}



func InSliceIface(ele interface{}, sli []interface{}) bool {
	for _, v := range sli {
		if v == ele {
			return true
		}
	}
	return false
}
