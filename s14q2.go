package main
import "fmt"
func remove(slice []int,id int)[]int{
return append(slice[:id],slice[id+1:]...)
}
func main(){
slice:=[]int{100,200,300,400,500}
fmt.Println(slice[0:5])
slice=append(slice,600,700)
fmt.Println("after appending",slice)

slice=remove(slice,2)
fmt.Println("after remove",slice)

newSlice:=make([]int,len(slice))
copy(newSlice,slice)
fmt.Println("newslice after",newSlice)
}
