package main
import "fmt"
type book struct{
bid int
title,author string
price float64
}
func main(){
var b[10]book
var n int
fmt.Println("enter the number of book")
fmt.Scanln(&n)
fmt.Println("enter the deatils of books")
for i:=0;i<n;i++{
fmt.Println("enter the bookid")
fmt.Scanln(&b[i].bid)
fmt.Println("enter the booktitle")
fmt.Scanln(&b[i].title)
fmt.Println("enter the bookauthor")
fmt.Scanln(&b[i].author)
fmt.Println("enter the bookprice")
fmt.Scanln(&b[i].price)
}
fmt.Println("detail of the book")
fmt.Println("bookid   title  author  price")
fmt.Println("-------------------------------")
for i:=0;i<n;i++{
fmt.Println(b[i].bid,"\t\t",b[i].title,"\t\t",b[i].author,"\t\t",b[i].price)
}
}
