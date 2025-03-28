package main
import "fmt"
func pal(n int){
m:=n
k:=0
for n!=0{
r:=n%10
k=k*10+r
n=n/10
}
if m==k{
fmt.Println("number is palindrom")
}else{
fmt.Println("number is not plaindrome")
}
}

func main(){
var n int
fmt.Println("enter the number")
fmt.Scanln(&n)
pal(n)
}
