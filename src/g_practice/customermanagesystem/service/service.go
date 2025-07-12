package service

import (
	//"fmt"
	"customermanagesystem/model"
)

type CustomerService struct{
	//用切片来存储用户信息
	customerlist []model.Customer
	//存储用户数量
	customerNum uint64
}
//工厂函数初始化service
//1.先定义一个字段来接收一个空结构体
//1.2对该结构体初始化字段，num和[]list
//1.3在初始化[]list时使用append方法，将customer（根据NewCustomer创建的一个客户信息）添加到切片[]list中
func NewCustomerService() *CustomerService{
	customerservice :=&CustomerService{}		
	customerservice.customerNum = 1
	customer :=model.NewCustomer(1,"张三","男",18,"12345678901","123456@qq.com")
	customerservice.customerlist =append(customerservice.customerlist,customer)
	return customerservice
}
//Question: *customerlist->*[]model.Customer
// 返回值类型必须为具体的像int []int等等
func (ser *CustomerService) List() *[]model.Customer{
	return &ser.customerlist
}
func (ser *CustomerService) AddCustomer(customer model.Customer) bool{
	ser.customerNum++
	customer.Id=ser.customerNum
	ser.customerlist  =append(ser.customerlist,customer)
	return true
}

func (ser *CustomerService) DelCustomer(name string) bool{
	index1:=true
	for index,val :=range ser.customerlist{
		if val.Name==name{
			ser.customerlist =append(ser.customerlist[:index],ser.customerlist[index+1:]...)
			ser.customerNum--
			index1=false
			break
		}
		}
		return !index1
	}
	
func (ser *CustomerService) ReCustomer(name string,
				chance string,args ...interface{}) bool{
	for i,_ :=range ser.customerlist{
		if ser.customerlist[i].Name==name{
			switch chance{
			case "性别":
				if v,ok :=args[0].(string);ok{
				ser.customerlist[i].Gender=v
				return true}
			case "年龄":
				 if v,ok :=args[0].(uint8);ok{
				ser.customerlist[i].Age=uint8(v)
				return true}
			case "电话":
				 if v,ok :=args[0].(string);ok{
				ser.customerlist[i].Phone=v
				return true}
			case "电邮":
				 if v,ok :=args[0].(string);ok{
				ser.customerlist[i].Gmail=v
				return true	}	
			}
			
		}
	
		}
		return false
	}	