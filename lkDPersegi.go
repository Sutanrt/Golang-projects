package main 

import ("fmt";"math")

func main () {
	var panjang,lebar,luas,keliling int
	var diagonal float64
	fmt.Scan(&panjang,&lebar)
	luas=panjang*lebar
	keliling=2*(panjang+lebar)
	diagonal=math.Sqrt(float64(panjang*panjang+lebar*lebar))
	fmt.Print("Luas:",luas,"\nKeliling:",keliling,"\nDiagonal:",diagonal)

} 