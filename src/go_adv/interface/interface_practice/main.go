package main
import (
	"fmt"
	"sort"
	"math/rand"
)

type Student struct{
	Name string
	Age int
	Score float64
}
//使用Sort里的对结构体进行sort的函数，需要先实现该方法传参的接口的三个方法
type Interface interface{
	Len() int
	Less(i ,j  int) bool
	Swap(i , j int)
}
type Stu []Student
func (s Stu) Len() int{
	return len(s)
}
//升序排列
func (s Stu) Less(i,j int) bool{
	return s[i].Score<s[j].Score
}

func (s Stu) Swap(i, j int){
		s[i],s[j]=s[j],s[i]
}

func main(){
	//接口实现完，开始验证
//1.先给切片随机填充数据
	var stus Stu//上面只是给Student切片换了个名字，并没有声明一个变量接收该切片
	for i:=0;i<10;i++{
		//stu1是切片中的任意元素，通过一个例子一直循环来进行每次添加新的Student
		stu1:=Student{
			//因为Sprintf会返回一个字符串，正好弥补了：后一起添加字符串和数字的问题
			Name:	fmt.Sprint("学生|",rand.Intn(100),),
			Age:	rand.Intn(100),
			Score:	rand.Float64()*100,
		}
		stus=append(stus,stu1)
	}

sort.Sort(stus)

	for _,v:= range stus{
		fmt.Println(v)
	}

}
