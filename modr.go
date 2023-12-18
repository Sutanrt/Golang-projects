package main 
import "fmt"


func main () {
 var n,a,b,c int
	a=1
	b=0
	c=1
 fmt.Scan(&n)
 procFibonacci(n,&a,&b,&c)

}

func procFibonacci (n int,a,b,c *int) {
	if n ==-1 {
		*a= 1
		//fmt.Print("0 ",*a)
	
	}else{
		fmt.Print(*b," ")
		*c=*b
		*b=*b+*a
		*a=*c
		
		//fmt.Println(*a,*b,*c)
		procFibonacci(n-1,a,b,c)
		
	}
	

}