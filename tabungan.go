package main 
import "fmt"
func main ()  {
	var tabungan,tMax,tMin,jumlahTabungan int
	var rata,n float64
	tMax=-1
	n=-1
	tMin=1000000000
	jumlahTabungan=0
	for tabungan>=0{
		jumlahTabungan=jumlahTabungan+tabungan
		if tMax<tabungan{
			tMax=tabungan
		}
		
		n=n+1
		fmt.Scan(&tabungan)
		if tMin>tabungan && tabungan>=0{
			tMin=tabungan
		}

	}
	rata=float64(jumlahTabungan)/n
	
	fmt.Println("Jumlah = ",jumlahTabungan)
	fmt.Println("Rata-rata= ",rata)
	fmt.Println("Uang tertinggi = ",tMax)
	fmt.Print("Uang terendah = ",tMin)

} 