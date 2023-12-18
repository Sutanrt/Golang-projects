package main 
import "fmt"

func recursive (n float64) float64  {

	if n==1 {
		return 1
	}else {
		return 1/n+recursive(n-1)
	}

}

func main () {
	var n float64
	fmt.Scan(&n)
	fmt.Print(recursive(n))

}