package main
import "fmt"
func jumlahBus_1301223465(penumpang,kapasitas int) int {
	var jumlah int
	jumlah =penumpang/kapasitas
	if penumpang%kapasitas!=0{
		jumlah=jumlah+1
	}
	return (jumlah)
} 
func main ()  {
	var penumpang,kapasitas int
	fmt.Scan(&penumpang,&kapasitas)
	fmt.Print(jumlahBus_1301223465(penumpang,kapasitas)," ","bus")

}