package main

import (
	"fmt"
	"customermanagesystem/service"
	"customermanagesystem/model"
)

type CustomerView struct{
	//1.用于接收用户输入
	key string
	//2.用于判断何时退出循环
	loop bool
	//3.创建service层对象,便于直接调用service层方法
	customerService *service.CustomerService
	
}

func (this *CustomerView) MainView(){
	for {
			fmt.Println("-----------------客户信息管理软件-----------------")
			fmt.Println("				  1 添 加 客 户")
			fmt.Println("				  2 修 改 客 户")
			fmt.Println(" 				  3 删 除 客 户")
			fmt.Println("				  4 客 户 列 表")
			fmt.Println(" 				  5 退 出")
			fmt.Print("请选择(1-5)：")
			fmt.Scanln(&this.key)
			switch this.key {
			case "1" :
				this.AddView()
			case "2" :
				this.ReView()
			case "3" :
				this.DelView()
			case "4" :
				this.list()
			case "5" :
				this.loop = false
		}
		if !this.loop{
			break
		}
	}

}

//显示客户列表
func (this *CustomerView) list(){
	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t \t邮箱")
	for i := 0; i < len(*customers); i++ {
	//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
	fmt.Println((*customers)[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
	}

	//添加界面
func (this *CustomerView) AddView(){
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//创建一个客服对象接收输入的信息
	customer := model.NewCustomer2(name, gender, uint8(age), phone, email)
	if this.customerService.AddCustomer(customer){{
		fmt.Println("添加成功")
	}}else
	{
		fmt.Println("添加失败")
	}
}

func (this *CustomerView) DelView(){
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)

	if this.customerService.DelCustomer(name){{
		fmt.Println("删除成功")
	}}else
	{
		fmt.Println("删除失败")
	}
}

func (this *CustomerView) ReView(){
	fmt.Println("---------------------修改客户信息---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入你想修改的内容:性别，年龄，电话，电邮")
	chance :=""
	fmt.Scanln(&chance)
		var found bool
	switch chance{
		case "性别":
			fmt.Println("请输入性别:") 
			gender := ""
			fmt.Scanln(&gender)
			found = this.customerService.ReCustomer(name,chance,gender)
		case "年龄":
			fmt.Println("请输入年龄:") 
			var age uint8= 0
			fmt.Scanln(&age)
			found =this.customerService.ReCustomer(name,chance,age)
		case "电话":
			fmt.Println("请输入电话:") 
			phone := ""
			fmt.Scanln(&phone)
			found =this.customerService.ReCustomer(name,chance,phone)
		case "电邮":
			fmt.Println("请输入电邮:") 
			email := ""
			fmt.Scanln(&email)
			found =this.customerService.ReCustomer(name,chance,email)

	}
	if found{
		fmt.Println("修改成功")
	}else{
		fmt.Println("修改失败")

	}

}
func main(){
	//创建一个CustomerView对象,直接调用方法和字段
	mainview :=CustomerView{
		loop: true,
		key : "",
	}
	mainview.customerService = service.NewCustomerService()
	mainview.MainView()
	

	
}