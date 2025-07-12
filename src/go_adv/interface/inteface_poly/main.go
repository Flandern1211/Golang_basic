package main

import "fmt"


func main(){
	var x interface{}
	var a float32 =1.1
	x=a//a是包含在x中的一个实现，所以可以直接让x=a，大化小可以，小化大永远不行，但可以是使用类型断言去获得 被大包含的小
	//b:=x//会报错，因为x其中可以包含多个实现，不指定是哪个实现类型无法执行
	// b:=x.(float32)//这样就可以了
	// fmt.Println(b)
	b:=x
	//可以加一个判断机制来阻止类型转换失败后程序直接退出，不再往下进行的机制
	b,ok:=x.(float64)
	if ok{
		fmt.Println("类型转换成功")
		fmt.Printf("b的值为%v",b)
	}else{
		fmt.Println("转换失败")
	}
	fmt.Println("结束")
}