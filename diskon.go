package main 
import "fmt"

func diskon (a int,m bool) float64{
	var disc float64
	disc=0
	if a>100000&&m{
		disc=0.1
	
	}else if a>100000{
		disc=0.05
	}
	//fmt.Println(a>100000,m)
	return disc

}

func main () {
	var belanja int 
	var disc float64
	var member bool
	fmt.Scan(&belanja,&member)
	disc=diskon(belanja,member)
	fmt.Print(float64(belanja)-(float64(belanja)*disc))

}