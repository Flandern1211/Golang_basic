package main

import "fmt"

type Calcuator struct{

}

func (c Calcuator) cal(num1 int , num2 int ,str byte) int{
	switch str{
	case '+' :
		return num1+num2
	case '-':
		return num1-num2
	case '*':
		return num1*num2
	case '/':
		return num1/num2
	default:
		fmt.Println("计算符输入有误")
	}
	return 0
}

func main(){
	var ca Calcuator
	fmt.Println(ca.cal(15,2,'+'))
}