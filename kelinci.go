package main 
import "fmt"

type c struct {
	nilai int
	min int
	max int
}
type x [1000]c
func main () {
var n int
var i int
var x x
fmt.Scan(&n) 
for i<len(x)&& i<n{
	fmt.Scan(&x[i].nilai)
	i=i+1
	
}

minmax(&x,n)
fmt.Print(x[0].min,x[1].max)
}
func minmax(x *x,n int)  {
	var min = x[0].nilai
	var max =x[0].nilai
	var i int
	for i< n {
		if min>x[i].nilai{
			min=x[i].nilai
		}
		if max<x[i].nilai {
			max=x[i].nilai
		
		}
		i=i+1
	}
	x[0].min=min
	x[1].max=max
}

/*func minmax(x [1000]int,n int) (int,int) {
	var min = x[0]
	var max =x[0]
	var i int
	for i< n {
		if min>x[i]{
			min=x[i]
		}
		if max<x[i] {
			max=x[i]
		
		}
		i=i+1
	}
	return min ,max
}*/