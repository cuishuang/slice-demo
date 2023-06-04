package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var sli []int
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		//var mu sync.Mutex
		go func() {
			// var mu sync.Mutex
			for j := 0; j < 10; j++ {
				// var mu sync.Mutex
				mu.Lock()
				sli = append(sli, 1)
				mu.Unlock()
			}
		}()
	}

	time.Sleep(time.Microsecond)

	time.Sleep(5e9) // 不加这一行，结果一般小于100；100次加锁解锁操作，1ms内完不成
	fmt.Println(len(sli))

}
