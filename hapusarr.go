package main 
import "fmt"
const n int =10
func main() {
	var x [n]int
	var indxHapus int
	for i:=0;i<n;i++{
		fmt.Scan(&x[i])
	}
	fmt.Println(x)
	fmt.Scan(&indxHapus)
	hapusArray(&x,indxHapus)
	fmt.Println(x)
}

func hapusArray(x *[n]int,indxHapus int) {
	var i int
	for i=indxHapus;i<n-1;i++{
		x[i]=x[i+1]
	}
}