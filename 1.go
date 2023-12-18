package main 
import "fmt"
func abc_1301223465(n int, s string) string {
	if n==0{
		return s
	}else {
		if n%2==0{
			s=s+"0"
		}else {
			s=s+"1"
		}
		return abc_1301223465(n/2,s)
	}
	

}
func main () {
	var n int 
	var s string
	fmt.Scan(&n)
	fmt.Print(abc_1301223465(n,s))

}


/*package main 
import "fmt"
func abc_1301223465(n int, s string) string {
	if n==0{
		return s
	}else if n==1{
	 return "1"+s
	}else {
		if n%2==0{
			s="0"+s
		}else {
			s="1"+s
		}
		return abc_1301223465(n/2,s)
	}
	

}
func main () {
	var n int 
	var s string
	fmt.Scan(&n)
	fmt.Print(abc_1301223465(n,s))

}*/