package main
import "fmt"
// 	 *
//	***
// *****
//*******
func main(){
	//1.先确定层数
	var flat int
	fmt.Println("请输入打印的层数")
	fmt.Scanf("%d",&flat)
	//2.1先将层数可以表示出来
	for i:=1;i<=flat;i++{
		//2.2开始打印每一层的*
		//2.2.1在打印之前明白每一层都只打印*的话会只有一半，因为都是从最左边开始打印
		//2.2.2先打印空格来使得*可以合起来显示成一个三角形
		//2.2.3确定空格的打印
		for j:=0;j<flat-i;j++ {
			fmt.Print(" ")
		}
		for k:=1;k<=2*i-1;k++{
			if k==1 ||k==2*i-1 ||i==flat{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
	
}