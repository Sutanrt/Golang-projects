package main 
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	barisGanjil(n)

}
func barisGanjil(n int) {
	if n==1{
		fmt.Print(n," ")
	}else {
		barisGanjil(n-1)
		if n%2!=0{
			fmt.Print(n," ")
		}
	}
}