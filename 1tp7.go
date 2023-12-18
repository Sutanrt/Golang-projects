package main 
import "fmt"

func inputBilangan_1301223465( bil *int) {
	fmt.Scan(&*bil)
	for *bil <0 {
		fmt.Scan(&*bil)
	}
}
func stop_1301223465 (bil int) bool{
	return bil==0

}
func hitung_1301223465 (mean *float64,min,max,n *int) {
	var bil int
	*max = -1
	inputBilangan_1301223465(&bil)
	*min=bil
	for !stop_1301223465(bil) {
		*n=*n+1
		*mean = float64(bil)+*mean
		if *min>bil{
			*min=bil
		}
		if *max < bil {
			*max=bil
		}
		inputBilangan_1301223465(&bil)
	}
	if *n>0{
		*mean=*mean/float64(*n)
		
		 
	}
	
	
}
func main () {
	var mean float64
	var min, max , n int 
	hitung_1301223465(&mean,&min,&max,&n)
	if n>0 {
		fmt.Print(mean," ",max," ",min," ",n)
	}else {
		fmt.Print("- - - -")
	}

}