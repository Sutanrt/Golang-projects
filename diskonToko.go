package main 
import "fmt"
func main () {
	var tahun int
	var belanja,diskon float64
	fmt.Scan(&tahun,&belanja)
	diskon=float64((tahun/1000)*tahun%100)
	belanja=belanja-(belanja*diskon/100)
	fmt.Println("Besar diskon: ",diskon,"%")
	fmt.Printf("Jumlah yang dibayar: %v",belanja)
	
}