package main 
import "fmt"


type arrBalita [100]float64

func main () {
	var n,i int
	var min, max float64
	var x arrBalita
	fmt.Scan(&n)
	for i<len (x) && i < n {
		fmt.Scan(&x[i])
		i=i+1
	}
	fmt.Println("cek")
	min=x[0]
	max=x[0]
	hitungMinMax(x,&min,&max)
	fmt.Printf("%.2f %.2f\n",min,max)
	fmt.Printf("%.2f",rerata(x))
}
func hitungMinMax(arrBerat arrBalita, bMin,bMax *float64) {
	var i int
	for i<len(arrBerat) && arrBerat[i]!=0{
		if *bMin>arrBerat[i]{
			*bMin=arrBerat[i]
		}
		if *bMax<arrBerat[i] {
			*bMax=arrBerat[i]
		}
		i=i+1
	}
}
func rerata (arrBerat arrBalita) float64 {
	var i,sum float64
	for arrBerat[int(i)]!=0{
		sum=sum+arrBerat[int(i)]
		i=i+1
	}
	return sum/i
}