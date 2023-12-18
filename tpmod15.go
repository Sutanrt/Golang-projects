package main 
import "fmt"



type negara struct {
	nama string
	gDP [3]int

}
type arrNegara [10]negara

func main () {
	var	ngr arrNegara
	var temp negara
	var n,idx,thn int
	var nama string
	
	ngr[0].nama="Malaysia"
	ngr[1].nama="Brunei"
	ngr[2].nama="Singapore"
	ngr[3].nama="Vietnam"
	ngr[4].nama="Myanmar"
	ngr[5].nama="Thailand"
	ngr[6].nama="Cambodia"
	ngr[7].nama="Philippines"
	ngr[8].nama="Indonesia"
	ngr[9].nama="Laos"
	
	fmt.Scan(&nama,&thn)
	idx=urut_1301223465(ngr,nama)
	for  idx !=-1{
		if thn==2021 {
			fmt.Scan(&ngr[idx].gDP[0])
		}else if thn==2022{
			fmt.Scan(&ngr[idx].gDP[1])
		}else if thn==2023{
			fmt.Scan(&ngr[idx].gDP[2])
		}
		fmt.Scan(&nama,&thn)
		idx=urut_1301223465(ngr,nama)	
	}
	for i:=1;i<=9;i++{
		temp=ngr[i]
		n=i
		if thn==2021{
			for n>0 && temp.gDP[0]>ngr[n-1].gDP[0]{
				ngr[n]=ngr[n-1]
				n=n-1
			}
			
		}else if thn==2022{
			for n>0 && temp.gDP[1]>ngr[n-1].gDP[1]{
				ngr[n]=ngr[n-1]
				n=n-1
			}
			//fmt.Println()
			//fmt.Println(ngr[i-1].nama," ",ngr[i-1].gDP[0]," ",ngr[i-1].gDP[1-1]," ",ngr[i-1].gDP[2])
		}else if thn==2023{
			for n>0 && temp.gDP[2]>ngr[n-1].gDP[2]{
				ngr[n]=ngr[n-1]
				n=n-1
			}
		}
		ngr[n]=temp
		
	}
	for i:=0;i<=9;i++{
		fmt.Println(ngr[i].nama," ",ngr[i].gDP[0]," ",ngr[i].gDP[1]," ",ngr[i].gDP[2])
	}

}
func urut_1301223465 (ngr arrNegara,nama string) int {
	var cek bool=false
	var idx int = -1
	var n int
		for !cek && n<=9{
			//fmt.Println(nama==ngr[n].nama,nama,"	",ngr[n].nama)
			cek=nama==ngr[n].nama
			if cek {
				idx=n
			}
			n=n+1
		}
		return idx
}
/*
Brunei 2023 889
Singapore 2021 676
Cambodia 2021 881
Myanmar 2021 950
Singapore 2023 712
Philippines 2021 120
Malaysia 2022 282
Philippines 2023 120
Indonesia 2023 62
Laos 2022 497
*/
/*
Brunei 2023 889
Singapore 2021 676
Cambodia 2021 881
Myanmar 2021 950
Singapore 2023 712
Philippines 2021 120
Malaysia 2022 282
Philippines 2023 120
Indonesia 2023 62
Laos 2022 497
Vietnam 2021 29
Cambodia 2023 277
Brunei 2022 469
Myanmar 2023 480
Thailand 2023 393
Thailand 2021 361
Laos 2021 108
Laos 2023 31
Malaysia 2023 968
Indonesia 2021 543
Vietnam 2023 590
Cambodia 2022 293
Brunei 2021 333
Singapore 2022 82
Myanmar 2022 354
Malaysia 2021 476
Thailand 2022 803
Philippines 2022 913
Vietnam 2022 711
Indonesia 2022 879
*/