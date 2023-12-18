package main 
import "fmt"
const NMAX = 1000000
var data [NMAX]int
func main () {
	var n,k int
	
	fmt.Scan(&n,&k)
	isiArray_1301223465(n)
	c:=posisi(n,k)
	if c==-1{
		fmt.Print("TIDAK ADA")
	}else {
		fmt.Print(c)
	}
}

func isiArray_1301223465(n int) {
	for i:=0;i<n;i++{
		fmt.Scan(&data[i])
	}
}
func posisi (n,k int) int {
	var mid,r,l,indx int
	
	l=0
	r=n-1
	mid=(l+r)/2
	indx=-1
	for l<=r && indx==-1 {
		if k<data[mid]{
			r=mid-1
		}else if k>data[mid]{
			l=mid+1
		}else{
			indx=mid
		}
		mid=(l+r)/2
	}
	return indx
}