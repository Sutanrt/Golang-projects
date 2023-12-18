package main 
import "fmt"
type basoe struct 	{
	banyak int
	jumlah int
	iterasi int
}

func main () {
	var x basoe
	
	fmt.Scan(&x.banyak)
	cek(&x)
	fmt.Print(x.jumlah)
}

func cek (n *basoe) {
	for n.iterasi<n.banyak{
		if n.iterasi %6 ==0 || n.iterasi %9==0|| n.iterasi %15==0|| n.iterasi %20==0{
			n.jumlah=n.jumlah+1
		}
		//fmt.Println(n.iterasi)
		n.iterasi++
	}
}