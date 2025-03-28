package main
import "fmt"
func main(){
var m1 [3][3]int
var m2 [3][3]int
var r [3][3]int

fmt.Println("enter numbers in 1st matrix")
for i:=0;i<3;i++{
     for j:=0;j<3;j++{
     fmt.Printf("enter the number in[%d][%d]",i,j)
     fmt.Scanln(&m1[i][j])
     }
   }
fmt.Println("enter number in 2nd matrix")
for i:=0;i<3;i++{
    for j:=0;j<3;j++{
    fmt.Printf("enter the number [%d][%d]",i,j)
    fmt.Scanln(&m2[i][j])
    }
   }

for i:=0;i<3;i++{
    for j:=0;j<3;j++{
         for k:=0;k<3;k++{
             r[i][j]+=m1[i][k]*m2[k][j]
             }
         }
     }
fmt.Println("multioalication is")
for i:=0;i<3;i++{
    for j:=0;j<3;j++{
     fmt.Printf("\t%d",r[i][j])
     }
     fmt.Println()
   }
}
