package main 
import "fmt"


func recursive (n int) int {
	
	if n==1 {
		return 1
	}else if n%2==0 {
		return n/2*(recursive(n-1))
	}
	
	return n*recursive(n-1)
	
}
func main() {
	var n int
	fmt.Scan(&n)
	x:=recursive(n)
	fmt.Print(x)
}