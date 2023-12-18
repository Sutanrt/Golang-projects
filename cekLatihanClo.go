package main 
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	fibonacci(n,0,1)
	fmt.Println(fibonacci2(n))
}
func fibonacci(n,a,b int) {
	if n==0 {
		fmt.Println(a)
	}else {
		fmt.Print(a," ")
		fibonacci(n-1,b,b+a)
	}
}
func fibonacci2(n int) int {
	if n==0 {
		return n
	}else if n==1{
		return n
	}else {
		return fibonacci2(n-2)+fibonacci2(n-1)
	}
}
