package main
import "fmt"

func main(){
var f1,f2,f3,n int
f1=0
f2=1
f3=f1+f2
fmt.Println("enter the number")
fmt.Scanln(&n)
fmt.Println("fibonacci series")
fmt.Println(f1)
fmt.Println(f2)
for i:=3; i<=n; i++{
fmt.Println(f3)
f1=f2
f2=f3
f3=f1+f2
}
}
