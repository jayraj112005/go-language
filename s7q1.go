package main
import "fmt"
type stud struct{
name,class string
per float64
}
func (s *stud) show(){
fmt.Println("student name:",s.name)
fmt.Println("student class:",s.class)
fmt.Println("student per:",s.per)
}
func main(){
s:=&stud{
    name:"krush",
    class:"TYBCA",
    per:93.20,
}
s.show()
}
