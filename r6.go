package main 
import "fmt"

func recursive_1301223465 (n int,s *string) {
	if n==0{	
		fmt.Println(*s)
		*s=""
	}else{
		*s=*s+"*"
		recursive_1301223465(n-1,s)
	}

}


func main () {
	var n int
	var s string
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		recursive_1301223465(i,&s)
	
	}
	
	
}