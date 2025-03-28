package main
import"fmt"
func main(){
evensum:=0
oddsum:=0
for i:=1;i<=100;i++{
if i%2==0{
evensum=evensum+i
}else{
oddsum=oddsum+i
}
}
fmt.Println("sum of evensum:",evensum)
fmt.Println("sum of oddno:",oddsum)
}
