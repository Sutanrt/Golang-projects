package main 
import "fmt"
func main () {
	var x,n int
	var cek bool
	cek=false
	fmt.Scan(&x,&n)
	for n>0 {
		cek=cek||(x==n%10)
		n=n/10
	}
	fmt.Print(cek)
	
}