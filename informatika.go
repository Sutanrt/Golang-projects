package main 
import ("fmt")

const NMAX int =12345
type prodi struct {
	nama, akreditasi string
	tahun,aktif,lulusan int
}
type fakultas struct {
	nama string
	arrProdi [NMAX-1]prodi
	N int
}

func main () {
	var fif fakultas
	var c prodi
	buatFakultas_1301223465(&fif)
	for j:=0;j<=9;j++{
		daftarProdi_1301223465(&fif,j)
		
	}
	c=terbanyak_1301223465(fif)
	fmt.Printf("Prodi %s memiliki mahasiswa dan lulusan terbanyak yaitu %d \n",c.nama,c.lulusan+c.aktif)
	
	fmt.Println("Prodi termuda adalah ",fif.arrProdi[termuda_1301223465(fif)].nama)
	fmt.Printf("Akreditasi Prodi terbanyak di Fakultas %s adalah %s\n",fif.nama,prestasi_1301223465(fif))
	cetakProdi_1301223465(fif,prestasi_1301223465(fif))
}
func buatFakultas_1301223465(f *fakultas) {
	fmt.Scan(&f.nama)
	
}
func daftarProdi_1301223465(f *fakultas, j int) {
	var nama string
	var i int
	var p prodi

		fmt.Scan(&nama)
			i=cekProdi_1301223465(nama,*f)
		if cekProdi_1301223465(nama,*f)==-1 {
			f.arrProdi[j].nama=nama
			fmt.Scan(&f.arrProdi[j].akreditasi,&f.arrProdi[j].tahun,&f.arrProdi[j].aktif,&f.arrProdi[j].lulusan)
			f.N=f.N+1
		}else {
			fmt.Scan(&p.akreditasi,&p.tahun,&p.aktif,&p.lulusan)
 
			f.arrProdi[i].aktif=f.arrProdi[i].aktif+p.aktif
			f.arrProdi[i].lulusan=f.arrProdi[i].lulusan+p.lulusan
		}
	
	

	
}

func cekProdi_1301223465(s string, f fakultas) int {
	var j int
	for j<f.N && f.arrProdi[j].nama!=s {
		j=j+1
	}
	if j==f.N {
		return -1
	}else {
		return j
	}
}
 func terbanyak_1301223465(f fakultas) prodi {
	var max int =-1
	var c prodi
	for i:=0 ; i<=9;i++{
		if max < f.arrProdi[i].aktif+f.arrProdi[i].lulusan {
			max=f.arrProdi[i].aktif+f.arrProdi[i].lulusan
			c.nama=f.arrProdi[i].nama
			c.aktif=f.arrProdi[i].aktif
			c.lulusan=f.arrProdi[i].lulusan
		}
	}
	return c
 }
 
 func termuda_1301223465 (f fakultas) int {
	var n int = -1
	var x int
	for i:=0 ; i<=9;i++{
		if n <= f.arrProdi[i].tahun {
			n=f.arrProdi[i].tahun
			x=i
		
		}
	}
	
	return x
 }
func prestasi_1301223465 (f fakultas) string {
	var unggul, baik, cukup ,x int
	var terbesar string
		for i:=0 ; i<=9;i++{
		if f.arrProdi[i].akreditasi=="unggul"{
			unggul=unggul+1
		}else if f.arrProdi[i].akreditasi=="baik"{
			baik = baik+1
		}else if f.arrProdi[i].akreditasi=="cukup"{
			cukup=cukup+1
		}
	}
	if baik> cukup {
		terbesar="baik"
		x=baik
	}else {
		terbesar="cukup"
		x=cukup
	}
	if x< unggul{
		terbesar="unggul"
	}
	return terbesar
}
func cetakProdi_1301223465(f fakultas, x string) {
	for i:=0 ; i<=9;i++{
		if x== f.arrProdi[i].akreditasi {
			fmt.Print(f.arrProdi[i].nama," ")
		}
	}
}