//实现将一个Monster对象序列化后保存在文件中
package testingcase02

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"io"
	
)

//1.结构体
type Monster struct {
	Name  string
	Age   int
	Skill []string
}

//2.定义其序列化方法
func (m *Monster) Store(filepath string,) (error){
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件时发生错误%v =",err)
		return  fmt.Errorf("打开文件时发生错误%v =",err)//细节把控
	}

	defer file.Close()

	//3.3创建一个基于file的缓冲区
	written := bufio.NewWriter(file)
	//Question: 因为这里不需要像read那样一行一行读取，直接就是一个字节切片，所以不需要用for循环
	data, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("序列化失败%v",err)
	}
	if _, err := written.Write(data);err != nil {
		return fmt.Errorf("写入失败%v",err)
	}
	if err := written.Flush(); err != nil {
		return fmt.Errorf("刷新缓冲区失败%v",err)
	}
	return nil
	
}

//对文件进行反序列化
func (m *Monster) ReStore(filepath string) (*Monster,error) {
	
	file,err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("文件打开失败%v",err)
	}
	defer file.Close()


	data,err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("文件读取失败%v",err)
	}

	if err := json.Unmarshal(data,&m); err !=nil{
		return nil,fmt.Errorf("反序列化失败%v",err)
	}
	return m,nil

}
func main() {
	//3.先创建一个个结构体实例
	m := Monster{
		Name: "牛魔王",
		Age : 12,
		Skill : []string{"jjj","ssss"},
	}
	//3.2进行对文件的操作
	filePath := "C:/Users/31800/Desktop/test.txt"
	mon,_:= m.ReStore(filePath)
	
		  
	
}
