package main//表明在main包中，每个go文件必须有归属的包

import (
    "fmt"
    "function01/test"
)//引用程序需要的包 为了使用包下的函数

func main(){
	
	fmt.Println(test.Studoo)
}