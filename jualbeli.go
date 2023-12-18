package main 
import "fmt"

type dataBarang struct {
	namaBarang string
	kategori string
	kataKunci string
}

type transaksiPenjualan struct{
	barangTerjual int
	tBarangDibeli int
	biayasatuBarang int
	barang dataBarang
}
type arrJualBeli [1000]transaksiPenjualan

func main () {
	var transaksi arrJualBeli
	var n int

}

func inputDataBarang (transaksi *arrJualBeli,n *int){
	var bnyk int
	fmt.Scan(&bnyk)
	for i:=0;i<bnyk;i++{
		fmt.Scan(&t[i].barang.namaBarang,&t[i].barang.kategori,&t[i].barang.kataKunci)
	}

}
func inputDataPenjualan (transaksi *arrJualBeli,n *int){
	var bnyk int
	fmt.Scan(&bnyk)
	for i:=0;i<bnyk;i++{
		fmt.Scan(&t[i].barangTerjual,&t[i].tBarangDibeli,&t[i].biayasatuBarang)
	}

}
func cariBarang ()
