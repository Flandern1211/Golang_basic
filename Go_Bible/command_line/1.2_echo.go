//用于给参数之间添加空格
package main

import (
	"fmt"
	"os"
)
func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]  // 第一次sep是空字符串，后续都是空格
        sep = " "              // 从第二个参数开始添加分隔符
    }
    fmt.Println(s)
}