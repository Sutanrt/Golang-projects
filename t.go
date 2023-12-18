package main
import "fmt"

func t (x int) string {
	var a int
	var s string
	s=""
	
	for x!= 0 {
		a=x%2
		x=x/2
		if a==0{
			s=s+"0"
			
		}else {
			s=s+"1"
		}
	} 
	return s
}
func main () {
	var x int
	fmt.Scan(&x) 
	fmt.Print(t(x))
}