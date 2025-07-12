package main
import "fmt"

// type Point struct{
// 	arr []int
// 	name string
// } 
type Point struct{
	x int
	y int 
}

type Rect struct{
	leftpoint,rightpoint Point
}
func main(){
	//point :=Point{[]int{1,2,3},"juju"}
	i:=Rect{Point{1,2},Point{3,4}}
}