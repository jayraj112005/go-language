package main
import "fmt"
func addsub(a,b int)(int,int,int){
sum:=a+b
diff:=a-b
mult:=a*b
return sum,diff,mult
}
func main(){
var n1,n2 int
fmt.Println("enter 1st no")
fmt.Scanln(&n1)
fmt.Println("enter 2nd no")
fmt.Scanln(&n2)
sum,diff,mult:=addsub(n1,n2)
fmt.Println("sum:",sum)
fmt.Println("diff:",diff)
fmt.Println("mult:",mult)
}
