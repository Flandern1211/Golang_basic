package main

import "fmt"

//1.首先先确定这个简单的软件的实现方式，用面向对象实现是可以的
//1.1确定使用面向对象实现，那结构体里面应该都有什么属性和方法
type HomeMoney struct{
	money uint64
	record []int64
}

//Question:少了初始函数
func NewHomeMoney() HomeMoney{
	return HomeMoney{
		money: 1200,
		record : []int64{100,-12,12,232},
	}
}

//3.1收支明细的实现：放在一个切片里，收入为正，支出为负，show是直接遍历切片来根据正负数来分类
//Question:方法都要用引用传递好一点
func (h *HomeMoney) ShowMoney(){
	fmt.Println("\t\t收支明细")
	for _,val :=range h.record{
		if val>0 {
			fmt.Println("收入",val)
		}else if val<0 {
			fmt.Println("支出",-val)
		}else{
			fmt.Println("未进行收入支出操作")
		}
	}
	fmt.Println("现在的总金额为",h.money)
}

//3.2
func (h *HomeMoney)ReIncome(){
	var income uint64
	fmt.Print("请登记收入: ")
	fmt.Scanf("%d", &income)
	_, _ = fmt.Scanln()
	h.record=append(h.record,int64(income))
	h.money+=income
		
}

func (h *HomeMoney)ReExpend(){
	var expend uint64
	fmt.Print("请登记支出: ")
	fmt.Scanf("%d", &expend)
	_, _ = fmt.Scanln()
	if expend>h.money {
		fmt.Println("钱不够")
		return
	}else{
		h.record=append(h.record,-int64(expend))//Question:expend类型转换
		h.money-=expend
	}

}
func main(){
	var chance uint8
	//Question:结构体定义不对: homeMonry HomeMoney
	homeMoney :=NewHomeMoney()
		
	//2.1显示出来简单的界面
	label1:
	for{
		fmt.Println("----------------家庭收支记账软件----------------\n")
		fmt.Println("\t\t1 收支明细\t\t")
		fmt.Println("\t\t2 登记收入\t\t")
		fmt.Println("\t\t3 登记支出\t\t")
		fmt.Println("\t\t4 退出\t\t\n")
		//Question:fmt.Scanf("请选择(1-4)%d",&chance)错误，
		fmt.Print("请选择(1-4): ")
		fmt.Scanf("%d", &chance)
	//Question:一直多循环了好几次
		_, _ = fmt.Scanln()
		//2.2进行选择判断来
		switch chance{
		case 1:
			homeMoney.ShowMoney()
		case 2:
			homeMoney.ReIncome()
		case 3:
			homeMoney.ReExpend()
		case 4:
			fmt.Println("程序已退出")
			break label1
	
		default:
			fmt.Println("请重新选择")
		}
	}
}