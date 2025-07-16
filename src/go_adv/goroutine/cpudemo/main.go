//设置cpu的使用数量
package main

import (
	"fmt"
	"runtime"
)

func main(){
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	//如果想预留几个cpu
	runtime.GOMAXPROCS(cpuNum-2)
	fmt.Println("ok")
}