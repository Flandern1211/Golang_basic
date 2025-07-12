package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	filePath := "C:/Users/31800/Desktop/test1.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("err of open %v", err)
		//加return直接退出函数,因为有defer中是在函数退出时才运行，如果文件读取失败，file.Close也会执行，但此时文件是空指针
		return

	}

	defer file.Close()

	readers := bufio.NewReader(file)

	for{
		str,err :=readers.ReadString('\n')
		if err !=nil{//读取到换行符
			if err == io.EOF && str != ""{
				// 处理最后一行没有换行符的情况
				fmt.Print(str)
			}
			break
		}
		fmt.Print(str) // 使用Print而非Println，避免重复添加换行符
	}

	writer := bufio.NewWriter(file)

	str := "hello,北京\n"

	for i := 0; i < 5; i++ {
		//Question:写入字符串的操作是WriterString
		writer.WriteString(str)
	}
	writer.Flush()

}