package main 
import "fmt"
func main () {
	
	var n int
	fmt.Scan(&n)
	var c=make([]int,n)
	var b=make([]bool,n+1)
	
	recursive(0,n,c,b)
}
func recursive (kd,n int,c [] int,b []bool) {
	if kd>=n {
		for i:=0;i<n;i++{
			fmt.Print(c[i])
			fmt.Println(c,kd,b)
		}
		fmt.Println()
		
	}else {
		
		for i:=1;i<=n;i++{
			if !b[i]{
				b[i]=true
				fmt.Println(c,kd,b)
				c[kd]=i
				recursive(kd+1,n,c,b)
				
				b[i]=false
				fmt.Println("basoe",c,b)
			
			
			
		}
	}
	
}
}


