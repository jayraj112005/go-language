package main
import "fmt"
func main(){
var n int
fmt.Println("enter a number")
fmt.Scanln(&n)
if n>=10 && n<=100{
fmt.Println("no is two digit")
}else{
fmt.Println("no is one digit")
}
}

