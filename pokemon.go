package main 
import "fmt"

type pokemon struct {
	nama string
	CP,HP,IVAtk,IVDef,IVHP int

}
type arrPokemon [2023]pokemon
func main () {
	var pkmn arrPokemon
	var n int
	var nama,flag string
	fmt.Scan(&n)
	inputPokemon(&pkmn,&n )
	fmt.Scan(&nama,&flag)
	printPokemon(pkmn,n)
	mengurutkanPokemon(&pkmn,n,flag)
	printPokemon(pkmn,n)
	fmt.Printf("Total IV %s adalah %.3f\n",nama,TotalIV(pkmn,n,nama))
}
func inputPokemon(T *arrPokemon,N *int) {
	for i:=0;i<*N;i++{
		fmt.Scan(&T[i].nama,&T[i].CP,&T[i].HP,&T[i].IVAtk,&T[i].IVDef,&T[i].IVHP)
	}


}
func printPokemon(T arrPokemon,N int){
	fmt.Println()
	for i:=0;i<N;i++{
		fmt.Printf("%s %d %d %d %d %d\n",T[i].nama,T[i].CP,T[i].HP,T[i].IVAtk,T[i].IVDef,T[i].IVHP)
	}

	//fmt.Printf("%s %d %d %d %d %d",T[i].nama,"	",T[i].CP,"	",T[i].HP,T[i].IVAtk," ",T[i].IVDef," ",T[i].IVHP)

}
func mengurutkanPokemon(T *arrPokemon, N int, flag string) {
	var j int
	var temp pokemon
	if flag=="CP"{
		for i:=1;i<N;i++{
			j=i
			temp=T[i]
			for j>=1 && temp.CP>T[j-1].CP{
				T[j]=T[j-1]
				j=j-1
			}
			T[j]=temp
		}
	}else if flag=="HP"{
		for i:=1;i<N;i++{
			j=i
			temp=T[i]
			for j>=1 && temp.HP>T[j-1].HP{
				T[j]=T[j-1]
				j=j-1
			}
			T[j]=temp
		}
	
	}else if flag=="IV_Atk"{
		for i:=1;i<N;i++{
			j=i
			temp=T[i]
			for j>=1 && temp.IVAtk>T[j-1].IVAtk{
				T[j]=T[j-1]
				j=j-1
			}
			T[j]=temp
		}
	}else if flag=="IV_Def"{
		for i:=1;i<N;i++{
			j=i
			temp=T[i]
			for j>=1 && temp.IVDef>T[j-1].IVDef{
				T[j]=T[j-1]
				j=j-1
			}
			T[j]=temp
		}
	}else if flag=="IV_HP"{
		for i:=1;i<N;i++{
			j=i
			temp=T[i]
			for j>=1 && temp.IVHP>T[j-1].IVHP{
				T[j]=T[j-1]
				j=j-1
			}
			T[j]=temp
		}
	}

}
func TotalIV(T arrPokemon, N int,namaX string) float64{
	var i int
	var jumlah int
	var index int =-1
	
	for i<N&&index==-1{
		if namaX==T[i].nama{
			index=i
		}
		i=i+1
	}
	if index==-1{
		return -9999
	}else{
		jumlah=T[index].IVAtk+T[index].IVDef+T[index].IVHP
		return (float64(jumlah)*100)/45
	}

}

/*
Note :
jika ingin konversi data dari int  ke float64 atau data tertentu lainnya 
lalu datanya banyak 
bisa digabung dulu ke suatu data tertentu misal jumlah 
yang jenis data nya sama dengan yang ingin diubah misal int
lalu baru ubah ke data yang ingin diuabh misal float64
hal itu lebih rapih dan mengurangi proses konversi menjadi lebih sedikit
*/
/*
6
Exxegutor 2733 173 15 13 14
Tyranitar 3753 182 15 14 6
Espeon 2791 131 12 13 13
Hippowdon 2784 190 13 15 13
Ho-Oh 2700 152 10 14 15
Slaking 3328 206 7 2 5
Pidgeot CP
Espeon IV_Def
*/