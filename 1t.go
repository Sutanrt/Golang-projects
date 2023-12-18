package main 

import "fmt"

func recursive (i,n int ) {
	if i<=n{
		fmt.Print(i," ")
		recursive(i+2,n)
	}




}




func main() {
 var n int
	i:=1
 fmt.Scan(&n)
	recursive(i,n)
}