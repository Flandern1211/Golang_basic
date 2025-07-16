// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

//查找有无重复行


//只有在不同行的重复行才行，使用ctrl +Z才算一个输出完成
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	//1.创建一个用于记录输入出现次数的map
    counts := make(map[string]int)
	//创建一个缓冲区用于保留用户的输入
    input := bufio.NewScanner(os.Stdin)
	//Scan()方法可以获取当前位置token，并让Scanner的扫描位置移动到下一个token。当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false
    for input.Scan() {
		//Text方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串。
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
	//遍历counts,因为键只有一个对应的值,其中只要有值大于1的就是重复的输入
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
