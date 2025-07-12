package main
import (
	"fmt"
	"go_adv/factory/model"
)

func main(){
	stu:=model.NewStudent("juju",12)
	fmt.Println(*stu)
	fmt.Println(stu.GetAge())
}