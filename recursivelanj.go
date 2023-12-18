package main 
import "fmt"

func recursive2(n int) int {
	if n==1 {
		fmt.Println("*")
		return 1
	}else {
		fmt.Print("*")
		return recursive2(n-1)
	}
}
func recursive(n int) int {
	if n ==1{
	 recursive2(n)
	 return 1
	}
	recursive(n-1)
	recursive2(n)
	recursive(n-1)
	return 1
	
}
func main () {
	var n int
	fmt.Scan(&n)
	recursive(n)
}