package main 
import "fmt"
func permutasian (x,y int)(int,int,int) {
	xfaktorial :=1
	yfaktorial:=1
	xyf:=1
	var permutasi int
	for i:=x;i>=1;i--{
		xfaktorial=xfaktorial*i
	}
	for i:=y;i>=1;i--{
		yfaktorial=yfaktorial*i
	
	}
	for i:=x-y;i>=1;i--{
		xyf=xyf*i
	
	}
	permutasi=xfaktorial/xyf
	return xfaktorial,yfaktorial,permutasi
}

func main () {
	var x,y ,xfaktorial,yfaktorial,permutasi int
	fmt.Scan(&x,&y)
	if x>y{
		xfaktorial,yfaktorial,permutasi=permutasian(x,y)
	}else{
		yfaktorial,xfaktorial,permutasi=permutasian(y,x)
	}
	fmt.Print(xfaktorial," ",yfaktorial," ",permutasi)

}