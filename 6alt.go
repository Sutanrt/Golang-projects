package main 
import "fmt"

func recursive (n int,s *string)  {
	if n==1{
		*s=*s+"*"
		fmt.Print(*s)
		
	
		
	}else{

		recursive(n-1,s)
		fmt.Println( )
		*s=*s+ "*"
		fmt.Print(*s)
	
	}


}

func main () {
	var n int 
	var s string
	
	fmt.Scan(&n)
	recursive(n,&s)
}