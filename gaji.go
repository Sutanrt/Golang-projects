package main 
import "fmt"
func main () {
	var jr,jl,tGaji int
	var g string 
	tGaji=0
	for g!="Z" {
		fmt.Scan(&g,&jr,&jl)
		if g=="A" {
			tGaji=tGaji+(jr*75000)+(jl*90000)
			fmt.Println((jr*75000)+(jl*90000))
		}else if g=="B" {
			tGaji=tGaji+(jr*125000)+(jl*70000)
			fmt.Println((jr*125000)+(jl*70000))
		}else if g!="Z"{
			fmt.Println("Golongan tidak dikenali")
		}
	
	}
	fmt.Print("Biaya yang dikeluarkan perusahaan untuk gaji karyawan ",tGaji)



}