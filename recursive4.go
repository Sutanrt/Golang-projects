package main 
import "fmt"
func recursive (n int) string {
	if n==1 {
		return "1"
	}
	
	
	return string(fmt.Sprint(recursive(n/2)+fmt.Sprint(n%2)))
}

func main () {
	var x int
	var s string
	fmt.Scan(&x)
	s=recursive(x)
	fmt.Print(s)
}