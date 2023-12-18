package main 
import "fmt"

func proc (n int){
	fmt.Print(n," ")
	for n>=1{
		if n==1 {
			fmt.Print()
			n=0
		
		}else if n%2==0 {
			n=n/2
			fmt.Print(n," ")
		}else if n%2!=0{
			n=(3*n)+1
			fmt.Print(n," ")
		}
	
	}
	
}
func main () {
	var n int
	fmt.Scan(&n)
	proc(n)
}