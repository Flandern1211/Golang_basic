package main

import "fmt"
type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}

func (p Phone) Start(){
	fmt.Println("手机开始工作")

}
func (p Phone) Stop(){
	fmt.Println("手机结束工作")
}

type Cammera struct{

}
func (c Cammera) Start(){
	fmt.Println("相机开始工作")

}
func (c Cammera) Stop(){
	fmt.Println("相机结束工作")
}

type Computer struct{

}

func (c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}

func main(){
	computer:=Computer{}
	phone := Phone{}
	cammera :=Cammera{}

	computer.Working(phone)
	computer.Working(cammera)
}