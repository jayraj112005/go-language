package main
import (
        "fmt"
        "sort"
        )
        
func main(){
var n int
fmt.Println("enter array size")
fmt.Scanln(&n)
a:=make([]int,n)
fmt.Println("enter number in array")
for i:=0;i<n;i++{
fmt.Scanln(&a[i])
}
fmt.Println("before sortin array",a)
sort.Ints(a)
fmt.Println("after sorting array",a)
}
