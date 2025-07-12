package main

import (
	"fmt"
	"os" // 提供操作系统相关功能，如文件操作
	"bufio"// 提供缓冲IO功能
	"io"// 提供基本IO接口
)

func main(){
	// 文件打开用os.Open 记住用:=号因为函数返回两个值，用:=号自动识别
	file,err :=os.Open("C:/Users/31800/Desktop/test.txt")
	if err!=nil {
		fmt.Println("err of open",err)
		return
	}
	
	// defer语句必须在file变量声明后定义
	// defer语句：延迟执行函数调用，确保文件在使用后关闭
	// 此处使用匿名函数包装file.Close()，以便处理关闭文件时可能出现的错误
	defer func(){
		err :=file.Close()
		if err !=nil {
			fmt.Println("err of close",err)
		}
	}()

	readers := bufio.NewReader(file)

	for{
		str,err :=readers.ReadString('\n')
		if err != nil {
			if err == io.EOF && str != ""{
				// 处理最后一行没有换行符的情况
				fmt.Print(str)
			}
			break
		}
		fmt.Print(str) // 使用Print而非Println，避免重复添加换行符
	}
	fmt.Println("读取结束")
}