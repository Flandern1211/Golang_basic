package main
import "fmt"

type MethodUtils struct{

}

func (m MethodUtils) print1(){
	for i:=0;i<10;i++{
		for j:=0;j<8;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (method MethodUtils) print2(m int, n int){
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (method MethodUtils) print3(len int ,width int) int{
	return len*width
}

func (method MethodUtils) print4(num int){
	if num%2==0{
		fmt.Println(num,"是偶数")
	}else{
		fmt.Println(num,"是奇数")
	}
}

func (m MethodUtils) print5(len int , width int ,str byte){

	for i:=0;i<len;i++{
		for j:=0;j<width;j++{
			fmt.Printf("%c",str)
		}
		fmt.Println()
	}
}
func main(){
	var method MethodUtils
	method.print1()
	fmt.Println("--------------------------------------------------")
	method.print2(1,2)
	fmt.Println("--------------------------------------------------")
	fmt.Println(method.print3(3,4))
	method.print4(1)
	method.print5(5,4,'-')
}