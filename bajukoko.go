package main 
import "fmt"

func membeliKain_1301223465(uangAwal float64, tUang,tPengeluaran *float64 ) {
	*tPengeluaran=*tPengeluaran+(uangAwal/4)
}
func membeliBenangJahit_1301223465(uangAwal float64,tUang,tPengeluaran *float64) {
	*tPengeluaran=*tPengeluaran+(uangAwal/20)
}
func membuatPolaBaju_1301223465 (uangAwal float64,tUang,tPengeluaran *float64) {
	*tPengeluaran=*tPengeluaran+(uangAwal/200)
}

func menjahitBaju_1301223465 (uangAwal float64,tUang,tPengeluaran *float64) {
	*tPengeluaran=*tPengeluaran+(uangAwal/200)
}
func mengemasBju_1301223465 (uangAwal float64,tUang,tPengeluaran *float64) { 
	*tPengeluaran=(*tPengeluaran)+(3*uangAwal/200)
}
func mendistribusikan_1301223465 (uangAwal float64,tUang,tPengeluaran,tPemasukan *float64) {
	*tPengeluaran=(*tPengeluaran)+((3*uangAwal)/200)
	*tPemasukan=*tPemasukan+(uangAwal/2)
}
func main () {
	var modalAwal,tUang,tPengeluaran,tPemasukan float64
	fmt.Scan(&modalAwal)
	membeliKain_1301223465(modalAwal , &tUang,&tPengeluaran  )
	membeliBenangJahit_1301223465(modalAwal, &tUang,&tPengeluaran  )
	membuatPolaBaju_1301223465(modalAwal , &tUang,&tPengeluaran  )
	membuatPolaBaju_1301223465(modalAwal , &tUang,&tPengeluaran  )
	menjahitBaju_1301223465(modalAwal, &tUang,&tPengeluaran  )
	menjahitBaju_1301223465(modalAwal , &tUang,&tPengeluaran  )
	menjahitBaju_1301223465(modalAwal, &tUang,&tPengeluaran  )
	menjahitBaju_1301223465(modalAwal , &tUang,&tPengeluaran  )
	mengemasBju_1301223465(modalAwal, &tUang,&tPengeluaran  )
	mengemasBju_1301223465(modalAwal , &tUang,&tPengeluaran  )
	mendistribusikan_1301223465(modalAwal , &tUang,&tPengeluaran,&tPemasukan  )
	fmt.Printf("%.f %.f %.f",tPengeluaran,tPemasukan,tUang-tPengeluaran+*tPemasukan	)
}
