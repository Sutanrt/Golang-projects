package main
import "fmt"

func main () {
	var bil int
	var hasil float64
	fmt.Scan(&bil)
	average_1301223465(bil,1,&hasil)
	fmt.Print(hasil)
}

func average_1301223465 (bil, i int, hasil *float64) {
	if bil <10 {
		*hasil=(*hasil+float64(bil))/float64(i)
	}else {
		i=i+1
		*hasil=*hasil+float64(bil%10)
		average_1301223465(bil/10,i,hasil)
	}
}