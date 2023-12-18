package main 
import ("fmt";"math")

func recursive (a,b,k,x int) int {
	//fmt.Println(a*x+b,a*x,b)
	
	if k == 1 {
		if a*x+b<0 {
			return -(a*x+b)
		}else {
			return (a*x+b)
		}
		
	}
	
	
		return int(math.Abs(float64(a*recursive(a,b,k-1,x)+b)))
	
	
	

}

func main () {
	var a,b,k,x int
	fmt.Scan(&a,&b,&k,&x)
	y:=recursive(a,b,k,x)
	fmt.Print(y)
}