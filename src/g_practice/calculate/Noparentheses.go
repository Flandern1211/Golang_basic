//实现无括号计算器
//1.先对整个算式（字符串）进行遍历，将其中的数字和计算符分隔开来放在两个切片（因为切片可以动态增删）中
//根据瞪眼法看出来每个计算符和其被计算数的下标索引是相同的
//下标索引相同就可以对照计算,遍历符号切片
//先计算乘除，遍历到有乘除号的索引直接对应找到其被计算数的下标,将被计算数与其后一位进行对应符号的计算
//计算得出结果，直接代替被计算数的位置，并将其后一位直接在切片中删除（通过切片拼接），符号切片同样
//确保没有乘除号时再对加减进行操作，加减进行操作时可以不用删除，可以在计算的时候将后一位置为0，对加减没有影响了就


package main

import "fmt"

func main(){
	//获取用户输入
	var res string
	var nums []int
	var fuhao []string
	var sum int
	fmt.Print("请输入无括号算式")
	fmt.Scanf("%s",&res)
	result :=[]byte(res)
	for _,val  :=range result{
		if val>='0' &&val<='9'{
			nums=append(nums,int(val-'0'))
		}else{
			fuhao=append(fuhao,string(val))
		}
	}
	
	//先计算乘除
	for i :=0;i<len(fuhao); {
		if fuhao[i]=="*" || fuhao[i]=="/"{
			if fuhao[i]=="*" {
				nums[i+1]=nums[i]*nums[i+1]
				nums=append(nums[:i],nums[i+1:]...)
				fuhao=append(fuhao[:i],fuhao[i+1:]...)
			}else{
				nums[i+1]=nums[i]/nums[i+1]
				nums=append(nums[:i],nums[i+1:]...)
				fuhao=append(fuhao[:i],fuhao[i+1:]...)
			}
			
		}else{
			i++
		}
		
	}

	//在计算加减
	sum =nums[0]
	for i ,val :=range fuhao{
		if val=="+"{	
			sum+=nums[i+1]
		}else{
			sum-=nums[i+1]
		}
	}
	fmt.Println("结果为",sum)
}
