package main
import "fmt"

func main(){
var choice,n1,n2 int
fmt.Println("choose arthmatic operation:")
fmt.Println("1.add(+)")
fmt.Println("2.sub(-)")
fmt.Println("3.mult(*)")
fmt.Println("4.div(/)")
fmt.Println("enter your choice;")
fmt.Scanln(&choice)

fmt.Println("enter 1st number")
fmt.Scanln(&n1)
fmt.Println("enter 2nd number")
fmt.Scanln(&n2)
switch choice{
case 1:
       fmt.Println("addtion is",n1+n2)
case 2:
       fmt.Println("sub is",n1-n2)
case 3:
       fmt.Println("mult is",n1*n2)
case 4:
       fmt.Println("div is",n1/n2)
}
} 
