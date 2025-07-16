//在第一版的基础上实现了多位数运算
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func SplitMore(s string,n ...string) ([]string, []string) {
	ops := []string{}
	//emp := []byte(s) 创建了一个新的字节切片，它是原始字符串 s 的独立副本。下面详细解释：
	temp := []byte(s)
	for i := 0; i < len(temp); i++ {
		//target := v[0]
		for _, v := range n {
			if len(v) != 1 {
			continue
		}
			if string(temp[i]) == v {
				ops = append(ops, v)
				temp[i] = '|'
			}
		}
	}
	nums := strings.Split(string(temp), "|")
	return ops,nums
}

func main() {
	//获取用户输入

	var res string
	var fuhao []string
	var sum int
	fmt.Print("请输入无括号算式")
	fmt.Scanf("%s", &res)
	// result := []byte(res)
	// for _, val := range result {
	// 	if val >= '0' && val <= '9' {
	// 		nums = append(nums, int(val-'0'))
	// 	} else {
	// 		fuhao = append(fuhao, string(val))
	// 	}
	fuhao,nums1 :=SplitMore(res,"*" , "+", "-", "/")
	

	// 转换字符串数字为int
	nums := make([]int, len(nums1))
	for i, num1 := range nums1 {
		num, _ := strconv.Atoi(num1)
		nums[i] = num
	}

	fmt.Println(nums)
	fmt.Println(fuhao)
	//先计算乘除
	for i := 0; i < len(fuhao); {
		if fuhao[i] == "*" || fuhao[i] == "/" {
			if fuhao[i] == "*" {
				nums[i+1] = nums[i] * nums[i+1]
				nums = append(nums[:i], nums[i+1:]...)
				fuhao = append(fuhao[:i], fuhao[i+1:]...)
			} else {
				if nums[i+1] !=0{
				nums[i+1] = nums[i] / nums[i+1]
				nums = append(nums[:i], nums[i+1:]...)
				fuhao = append(fuhao[:i], fuhao[i+1:]...)
				}
				fmt.Println("除数不能为0")
				 return 
			}

		} else {
			i++
		}

	}

	//在计算加减
	sum = nums[0]
	for i, val := range fuhao {
		if val == "+" {
			sum += nums[i+1]
		} else {
			sum -= nums[i+1]
		}
	}
	fmt.Println("结果为", sum)
}
