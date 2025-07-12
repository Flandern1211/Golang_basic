package main
import "fmt"

func main(){

}

func makeSuffix(suffix string) func(string) string{
	//感觉闭包中的环境中的变量就像全局变量一样
	return func(fname string) string{
		if !strings.HasSuffix(fname,suffix){
			return fname+suffix
		}else{
			return name
		}
	}
}