package main
import "fmt"

type AsDos_T struct {
	kode string
	nama string
	nim int
	mk string
	jumlah int
}
type TabAsDos_T = [2500]AsDos_T

func main () {
	var tabel TabAsDos_T
	var n int

	
	BacaAsdos_1301223465(&tabel,&n)
	for i:=0; i<n;i++ {
		if tabel[i].mk!="SUDAH" {
			CetakAsdos_1301123465(&tabel,n,tabel[i].mk)
		}
	}
	
	
}

func BacaAsdos_1301223465(tabel *TabAsDos_T, n *int) {
	var kode string
	*n=-1
	for kode!="XXX" {
		*n=*n+1
		fmt.Scan(&kode,&tabel[*n].nama,&tabel[*n].nim,&tabel[*n].mk,&tabel[*n].jumlah)
		tabel[*n].kode =kode
	}
	
}

func CetakAsdos_1301123465 (tabel *TabAsDos_T,n int, mk string ) {
	fmt.Println("Asisten Dosen Mata Kuliah : ",mk)
	for i:=0;i<n;i++{
		if tabel[i].mk==mk{
			fmt.Println(tabel[i].nama," ",tabel[i].kode)
			tabel[i].mk="SUDAH"
		}
	}
}