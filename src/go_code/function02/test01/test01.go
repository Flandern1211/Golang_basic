package main

import(
	"fmt"
	"strconv"
) 

func main(){
	var str="juju"//不可以修改已经定义好的字符串
	str ="wangwang"
	var num1 int =99
	str[0] = 'a'
	fmt.Println(str)
	str1=strconv.FormatInt(int64(num1),10)
}