package test06
import "fmt"
func sum(slice []byte) int{
	sum1:=0
	for index,key range slice{
		sum1 +=slice[index]
	} 
	return sum1
}
func main(){
	var str string
	var str2 []byte =make([]byte 10,len(str))
	fmt.Println("请输入一个算式")
	fmt.Scanf("%v",&str)
	for index,val range str{
		if str[index]=='*' {
			str[index]:= int(str[index-1])*int(str[index+1])
			str[index-1]='0'
			str[index+1]='0'	
		}
		if str[index]=='/' {
			str[index]:= int(str[index-1])/int(str[index+1])
			str[index-1]='0'
			str[index+1]='0'		
		}
	}
	for index1,key range str{
		if str[index]!='0'{
			str2.append(str[index])
		}
	}
	for index2,key1 range str2{
		if str2[index]=='+' {
			str2[index]= byte(int(str2[index-1])+int(str2[index+1]))
			str2[index-1]='0'
			str2[index+1]='0'	
		}
		if str[index]=='-' {
			str2[index]:= byte(int(str2[index-1])-int(str[index+1]))
			str2[index-1]='0'
			str2[index+1]='0'		
		}
	}
	fmt.Println(sum(str2))
}