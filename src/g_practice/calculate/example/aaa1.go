package main

import (
	"fmt"
	"strings"
)

// precedence 返回运算符的优先级
// + - 为 1，* / 为 2，其余（如括号）为 0
func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

// applyOp 对两个整数 a、b 执行 op 运算
// 仅支持四则运算：+ - * /
func applyOp(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	panic("invalid operator")
}

// evaluate 把不含空格的纯四则运算表达式字符串计算成整数
// 支持括号、多位数、*/ 优先于 +-
func evaluate(expr string) int {

	var nums []int // 数字栈
	var ops []rune // 运算符栈

	// readNum 从 expr 的 i 位置开始，连续读取一个整数
	// 因为 for 循环本身会 i++，所以读完要回退一位
	readNum := func(i *int) int {
		num := 0
		//获取计算使用的数字，单位数和多位数都可以
		//1.循环遍历整个字符串，找到其中单个字符在'0'-'9'之间的字符，然后对其后一位进行判断，如果也在该范围中就和前面的组成一个多位数，直到将这个多位数完全得到就会停住
		//1.1 *i <len(expr):保证索引到最后一个停住
		//1.2 num = num*10 + int(expr[*i]-'0') ，单位数直接就是将前面一位转为int型后
		for *i < len(expr) && expr[*i] >= '0' && expr[*i] <= '9' {
			num = num*10 + int(expr[*i]-'0')
			*i++
		}
		*i-- // 回退一位，主循环会再++
		return num
	}

	// 1.1逐个字符扫描表达式
	for i := 0; i < len(expr); i++ {
		ch := rune(expr[i]) // 当前字符

		switch {
		case ch == '(':
			// 左括号直接压入运算符栈
			ops = append(ops, ch)

		case ch == ')':
			// 右括号：不断弹出运算符并计算，直到遇到左括号
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				b := nums[len(nums)-1]
				nums = nums[:len(nums)-1]

				a := nums[len(nums)-1]
				nums = nums[:len(nums)-1]

				nums = append(nums, applyOp(a, b, op))
			}
			ops = ops[:len(ops)-1] // 弹出左括号 '('

		case ch >= '0' && ch <= '9':
			// 数字：读取完整整数并压入数字栈
			nums = append(nums, readNum(&i))

		case ch == '+' || ch == '-' || ch == '*' || ch == '/':
			// 运算符：先弹出栈里优先级 ≥ 当前运算符的全部运算
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				b := nums[len(nums)-1]
				nums = nums[:len(nums)-1]

				a := nums[len(nums)-1]
				nums = nums[:len(nums)-1]

				nums = append(nums, applyOp(a, b, op))
			}
			// 当前运算符入栈
			ops = append(ops, ch)
		}
	}

	// 表达式扫描完毕，处理剩余运算符
	for len(ops) > 0 {
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		nums = append(nums, applyOp(a, b, op))
	}

	// 最终栈里只剩一个数字，即为结果
	if nums == nil {
		return -1
	}
	return nums[0]
}

func main() {
	for {
		//1.定义接受输入的字符串
		var text string
		fmt.Print("请输入算数表达式：")
		// 用 Scanln 读取一行
		if _, err := fmt.Scanln(&text); err != nil {
			fmt.Println("读取输入失败:", err)
			return
		}
		//2.定义退出机制
		if text == "p" {
			break
		}
		//3.返回一个将输入前后端空白都去掉的字符串
		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("输入为空")
			return
		}
		//4。对输入进行计算处理并返回结果
		res := evaluate(text)
		if res < 0 {
			fmt.Println("输入内容不能为空")
		}
		fmt.Printf("计算结果：%d\n", res)
	}

}
