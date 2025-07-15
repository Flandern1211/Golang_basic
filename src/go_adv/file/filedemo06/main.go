//拷贝图片，音频等
//io.Copy的使用方法
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(desFile string, srcFile string) (written int64, err error) {
	srcfile, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	reader := bufio.NewReader(srcfile)
	defer srcfile.Close()

	desfile, err := os.OpenFile(desFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("writen file err%v", err)
		return
	}

	defer desfile.Close()
	writen := bufio.NewWriter(desfile)

	return io.Copy(writen, reader)

}

func main() {
	srcfile := "C:/Users/31800/Desktop/Snipaste_2025-05-31_10-32-39.png"
	desfile := "C:/Users/31800/Desktop/test.png"
	_, err := CopyFile(desfile, srcfile)
	if err != nil {
		fmt.Println("拷贝失败")
	} else {
		fmt.Println("拷贝完成")
	}

}
