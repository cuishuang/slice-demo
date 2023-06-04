package main

import (
	"fmt"
	"time"
)

func main() {

	sl := []int64{}

	_ = sl
	//fmt.Println(InSliceInt64(1,sl))

	//fmt.Println(time.Now().Format("20060102"))

	fmt.Println(time.Now().Unix())

	//fmt.Println(int64(time.Hour))
	fmt.Println(int64(time.Minute) * 60)

	fmt.Println(time.Now().Unix() + int64(time.Hour))

}

func InSliceInt64(i int64, sl []int64) bool {
	for _, vv := range sl {
		if vv == i {
			return true
		}
	}
	return false
}

func Remove(sli []interface{}, value interface{}) error {

	var newSli []interface{}
	for _, item := range sli {

		if item != value {
			newSli = append(newSli, item)
			if 1 == 1 {
				fmt.Println(111)
			}
		}
	}

	sli = newSli

	return nil

}
