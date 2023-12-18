package main 
import "fmt"

func fibonacci (n int) int {
	a:=0
	b:=0
	c:=1
	sum:=0
	for i:=1;i<=n;i++{
		sum=sum+b
		a=b
		b=b+c
		c=a
	}
	return sum
}

func main () {
	var n, sum int
	fmt.Scan(&n)
	sum = fibonacci(n)
	fmt.Print(sum)

}