package main
import "fmt"
func addsum(a,b int)(int,int,int,int){
return a+b,a-b,a*b,a/b
}
func main(){
var n1,n2 int
fmt.Println("enter 1st no")
fmt.Scanln(&n1)
fmt.Println("enter 2nd no")
fmt.Scanln(&n2)
sum,sub,mult,div:=addsum(n1,n2)
fmt.Println("add no is",sum)
fmt.Println("sub no is",sub)
fmt.Println("mult no is",mult)
fmt.Println("div no is",div)
}
