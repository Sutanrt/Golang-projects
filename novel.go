package main 
import "fmt"

func beliBuku_1301223465(n, m int ) {
	if n==m {
		fmt.Print("beli 1 buku , rak buku penuh")
	}else if n>m{
		
		fmt.Println("beli 1 buku , sisa ",n-m)
		beliBuku_1301223465(n,m+1)
	}
}
func main () {
	var n,m int
	fmt.Scan(&n,&m)
	fmt.Printf("Sisa rak kosong %d buku\n",n-m)
	beliBuku_1301223465(n,m+1 )
}