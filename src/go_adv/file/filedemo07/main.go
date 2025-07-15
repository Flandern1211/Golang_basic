package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	Chcount  int
	Numcount int
	Othcount int
	Spacount int
}

func Process(s string, c *Count) {
	for _, v := range s {
		switch {
		case v >= '0' && v <= '9':
			c.Numcount++
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			c.Chcount++
		case v == ' ' || v == '\t':
			c.Spacount++
		default:
			c.Othcount++
		}
	}
}
func main() {
	count := Count{}
	srcfile := "C:/Users/31800/Desktop/test1.txt"
	file, err := os.Open(srcfile)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	//开始统计
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF && str != "" {
				Process(str, &count)
				break
			}
			break
		}
		Process(str, &count)
	}
	fmt.Printf("统计结果:\n字母:%d\n数字:%d\n空格:%d\n其他:%d\n",
		count.Chcount, count.Numcount, count.Spacount, count.Othcount)
}
