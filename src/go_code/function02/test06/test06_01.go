package main
import (
	"fmt"
	"strconv"
)

func main(){
	var str string
	var nums []int
	var ops []byte
	fmt.Println("请输入一个算式")
	fmt.Scanf("%s",&str)

	//1
	for i:=0;i<len(str);i++{
		if str[i]>='0' && str[i]<='9'{
			//2
			num,_ := strconv.Atoi(string(str[i]))
			//3
			nums=append(nums,num)
		}else{
			//3
			ops=append(ops,str[i])
		}
	}
	for i:=0;i<len(ops);{
		if ops[i]=='*' ||ops[i]=='/'{
			
			if ops[i]=='*'{
				nums[i]=nums[i]*nums[i+1]
			}else{
				nums[i]=nums[i]/nums[i+1]
			}
			nums = append(nums[:i+1],nums[i+2:]...)
		ops= append(ops[:i],ops[i+1:]...)
		}else{
			i++
		}
		
	}
	res := nums[0]
	for i:=0;i<len(ops);i++{
		if(ops[i]=='+'){
			res+=nums[i+1]
		}else{
			res-=nums[i+1]
		}
		
	}
	fmt.Println(res)
}