package main
import "fmt"

type Student struct{
	Name string
	Age int
}
// 为 *Student 类型实现 String() 方法
func(stu *Student) String() string{
	return fmt.Sprintf("Name的值为%v,Age的值为%v",stu.Name,stu.Age)
}

func main(){
	var stu1 Student =Student{
		Name:"juju",
		Age:12,	
	}
	fmt.Println(&stu1)
}