package main
import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string
	Age int
}

var bytes =[...]byte{'小'}
func main(){
	niu :=Monster{"牛魔王",12}
	jsonStr,err:=json.Marshal(niu)
	if err!=nil {
		fmt.Println("json处理错误",err)
	}
	fmt.Println("jsonStr",jsonStr)
	fmt.Println("jsonStr",bytes)
	fmt.Printf("jsonstr:%c",bytes[0])

}