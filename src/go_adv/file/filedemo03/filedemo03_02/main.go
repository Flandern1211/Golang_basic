//覆盖之前的全部内容
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "C:/Users/31800/Desktop/test1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("err of open %v", err)
		//加return直接退出函数,因为有defer中是在函数退出时才运行，如果文件读取失败，file.Close也会执行，但此时文件是空指针
		return

	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	str := "hello,wangwang\n"

	for i := 0; i < 10; i++ {
		//Question:写入字符串的操作是WriterString
		writer.WriteString(str)
	}
	writer.Flush()

}
