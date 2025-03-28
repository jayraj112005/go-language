package main
import "fmt"
type emp1 struct{
eno,esal int
ename string
}
func main(){
var n int
fmt.Println("enter no of emp")
fmt.Scanln(&n)
emp:=make([]emp1,n)
for i:=0;i<n;i++{
fmt.Println("enter emp no")
fmt.Scanln(&emp[i].eno)
fmt.Println("enter emp ename")
fmt.Scanln(&emp[i].ename)
fmt.Println("enter emp esal")
fmt.Scanln(&emp[i].esal)
}
minsal:=emp[0].esal
for i:=1;i<n;i++{
if emp[i].esal<minsal{
minsal=emp[i].esal
}
}
fmt.Println("-----------------------------")
fmt.Println("emp having min sal")
for _,emp:=range emp{
if emp.esal==minsal{
fmt.Println("employee no is",emp.eno)
fmt.Println("emp name is",emp.ename)
fmt.Println("emp sal is",emp.esal)
}
}
}
