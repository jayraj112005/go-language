package main
import "fmt"
func abc(c chan int){
for i:=1;i<=10;i++{
c<-i
}
close(c)
}
func main(){
c:=make(chan int)
go abc(c)
for num:= range c{
fmt.Println(num)
}
fmt.Println("channel close succesfully")
}
