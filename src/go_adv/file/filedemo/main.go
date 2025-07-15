package main

import (
	"fmt"
	"os"//os:operate system
)

func main(){
	//文件打开用os.Open 记住用:=号因为函数返回两个值，用:=号自动识别
	//func Open(name string) (file *File, err error)
	file,err :=os.Open("C:/Users/31800/Desktop/test.txt")
	if err!=nil {
		fmt.Println("err of open",err)

	}
	//关闭文件
	err1 :=file.Close()
	if err1!=nil{
		fmt.Println("err of close",err1)
	}
	fmt.Println(file)
}
