package main

import (
	"fmt"
	
)

type Dog struct{
	Name string
	Age int
	Skills string
}
func test(s []int) {
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println("test函数内的")
	fmt.Println(len(s), cap(s))
}
func main() {
	//var arr = [3]int{1,2,3}
	// slice := []int{1, 2, 3}
	// slice2 := make([]int, 3, 3)
	// // str := "123"
	// // a ,_:= strconv.ParseInt(str,10,10)
	// // fmt.Printf("%T,%v,%v\n", a, a, str)
	// // for i:=0;i<len(arr);i++ {
	// // 	arr[i]+=1
	// // }
	// // fmt.Println(arr)
	// slice2 = append(slice2, slice...)
	// slice2 = append(slice2, 4)
	// test(slice2)
	// fmt.Println("main函数内的")
	// fmt.Println(len(slice2), cap(slice2))
	// dog := Dog{}
	// delete()
	
	//  a := new(int)
	//  fmt.Println(a,*a)

	var m map[int]int
	m = make(map[int]int)
	m[0]=4
	m[1]=1
	m[2]=2
	m[3]=3

	if v,ok := m[3];ok{
		fmt.Println(v)
	}else{
		fmt.Println(v)
	}

// 	if ok,a := i.(string);ok{
// 		fmt.Println(ok)
// 	}
// 	s := "1 2 3"
// 	a := strings.
//
}


