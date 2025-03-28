package main
import "fmt"
func swap(x *int,y *int){
temp:=*x
*x=*y
*y=temp
}
func main(){
var p,q int
fmt.Println("enter 1st number")
fmt.Scanln(&p)
fmt.Println("enter 2nd number")
fmt.Scanln(&q)
fmt.Println("before swap:",p,q)
swap(&p,&q)
fmt.Println("after swap:",p,q)
}
