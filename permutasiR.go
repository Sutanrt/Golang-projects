package main
import "fmt"
func recursive (kd,n int,c []int,b []bool) {
	if kd==n {
		for i:=0;i<n;i++{
			fmt.Print(c[i])
		}
	}else {
		for i:=1;i<=n;i++{
			if !b[i] {
				c[i]=i
				
				recursive(kd+1,n,c,b)
			}
		}
	}
}

func main () {
	var n int 
	fmt.Scan(&n)
	c:=make([]int,n)
	b:=bool([]bool,n)
	
	
}