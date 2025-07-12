- [在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用](#go函数作为变量)

## Go函数作为变量

在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用。

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    var f func(int, int) int
    f = add
    fmt.Println(f(3, 4)) // 输出 7
}
```