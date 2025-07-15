package main

import (
	"fmt"
	"encoding/json"
)

//1.对结构体的序列化
type Monster struct {
	Name string
	Age int
	Gender string
	Skills string
}

func testStruct() {
	monster := Monster{"牛魔王",2341,"雄","牛魔拳"}
	//使用&monster传地址
	//使用Marshal序列化       
	data , err := json.Marshal(&monster)
	if err !=nil {
		fmt.Printf("err from json%v\n",err)
		return
	}
	fmt.Println("序列化后的值为=",string(data))
}
//2.Map的序列化
func testMap(){
	var m map[string]interface{}
	//使用map要用make
	m =make(map[string]interface{})
	m["姓名"]="牛魔王"
	m["技能"] = [3]string{"牛魔拳","芭蕉扇","三味真火"}
	data , err := json.Marshal(m)
	if err !=nil {
		fmt.Printf("err from json%v\n",err)
		return 
	}
	fmt.Println("序列化后的值为=",string(data))
}

//3.切片的序列化
func testSlice(){
	var s []map[string]interface{}
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["姓名"] = "牛魔王"
	m1["技能"] = [3]string{"牛魔拳","芭蕉扇","三味真火"}

	m2 = make(map[string]interface{})
	m2["姓名"] = "琚振兴"
	m2["技能"] = [2]string{"王八拳","铁头功"}
	s = append(s,m1)
	s = append(s,m2)

	data , err := json.Marshal(s)
	if err !=nil {
		fmt.Printf("err from json%v\n",err)
		return
	}
	fmt.Println("序列化后的值为=",string(data))


}
func main(){
	testStruct()
	testMap()
	testSlice()
}