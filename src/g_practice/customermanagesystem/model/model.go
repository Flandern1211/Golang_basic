package model

import ("fmt"
	"strings"
)
type Customer struct{
	Id uint64
	Name string
	Gender string
	Age uint8
	Phone string
	Gmail string

}
//工厂模式
func NewCustomer(id uint64,name string,gender string,
		age uint8,phone string,gmail string) Customer{
			return Customer{
				Id : id,
				Name : name,
				Gender : gender,
				Age : age,
				Phone : phone,
				Gmail : gmail,
			}
		}

//用于返回显示客户列表
func (this1 *Customer) GetInfo() string{
	info := fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s\t",this1.Id,
		this1.Name,this1.Gender,this1.Age,this1.Phone,this1.Gmail)
	return info
}

//实现一个不带id号的新增客户实例化对象
func NewCustomer2(name string,gender string,
		age uint8,phone string,gmail string) Customer{
			return Customer{
				Name : name,
				Gender : gender,
				Age : age,
				Phone : phone,
				Gmail : gmail,
			}
		}