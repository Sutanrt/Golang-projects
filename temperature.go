package main 
import "fmt"

func main () {
	var x,tinggi,rendah,cek,jumlah int
	var rata,i float64
	for cek<=1 {
		fmt.Scan(&x)
		if x==0 {
			cek=cek+1
		}else{
			cek=0
		}
		
		jumlah=jumlah+x
		if tinggi<x{
			tinggi=x
		}
		if rendah>x{
			rendah=x
		}
		i=i+1
	
			
	}
	rata=float64(jumlah)/(i-1)
	fmt.Print(tinggi," ",rendah," ",rata)
	

}