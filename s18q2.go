package main
import "fmt"
func mult(n int){
for i:=1;i<=10;i++{
fmt.Println("table is",i*n)
}
}
func main(){
var n int
fmt.Println("enter the number of mult table")
fmt.Scanln(&n)
mult(n)
}
