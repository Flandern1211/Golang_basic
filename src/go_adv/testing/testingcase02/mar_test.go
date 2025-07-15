package testingcase02

import (
	"testing"
)

m := Monster{
		Name: "牛魔王",
		Age : 12,
		Skill : []string{"jjj","ssss"},
	}
	//3.2进行对文件的操作
	filePath := "C:/Users/31800/Desktop/test.txt"

func TestStore(t *testing.T){
	err := m.Store(filePath)
	if err != nil {
		t.Fatalf("序列化失败%v",err)
	}
	t.Logf("Store(filePath)执行正确")
}


func TestReStore(t *testing.T){
	_,err := m.ReStore(filePath)
	if err != nil {
		t.Fatalf("反序列化失败%v",err)
	}
	t.Logf("REStore(filePath)执行正确")
}