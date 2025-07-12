package main
import "fmt"

func fbn(n int) []uint{//定义时可以关注下范围
	if n<=0{
		return []uint{}
	}
	if n==1{
		return []uint{1}
	}
	var slice []uint = make([]uint,n,n+1)
	slice[0]=1
	slice[1]=1
	for i:=2;i<len(slice);i++{
		slice[i]=slice[i-1]+slice[i-2]
	}
	return slice
}

func main(){
	var n int
	fmt.Println("输入你想得到的数列数量")
	fmt.Scanf("%d",&n)
	slice:=fbn(n)
	fmt.Println(slice)
}