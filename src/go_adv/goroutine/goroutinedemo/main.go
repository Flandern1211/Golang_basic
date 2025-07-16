package main

import (
	"fmt"
	"time"
)

var (
	m :=make(map[int]int,10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	m[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//休眠十秒
	time.Sleep(time.Second*10)

	for i, v := range m {
		fmt.Printf("map[%d]=%v\n", i, v)
	}
}
