package main 
import "fmt"

type data struct {
	umur int
	nama string
	kelamin string
	golonganDarah rune

}
type aData [100]data

func main () {
	var arData aData
	var index,n,umurDicari,min,max int
	inputData(&arData,&n)
	min,max=carMaxMinUmur(arData,n)
	fmt.Println("halo")

	fmt.Scan(&umurDicari)

	fmt.Println("halo")
	fmt.Println(arData[min],arData[max])
	sortingUmur(&arData,n)
	fmt.Println(arData)
	index=cariUmur(arData,n,umurDicari)
	fmt.Println("halo")
	if index==-1{
		fmt.Print("Umur tidak ditemukan")
	}else{
		fmt.Println(arData[cariUmur(arData,n,umurDicari)])
	}
	
	
}
func sortingUmur (arData *aData,n int) {
	var temp,j int
	var temp1,temp2 string
	var temp3 rune
	fmt.Println(arData)
	for i:=1;i<n;i++{
		j=i
		fmt.Println(j,"H")
		for arData[j-1].umur>arData[j].umur && j>=1{
			temp=arData[j-1].umur
			temp1=arData[j-1].nama
			temp2=arData[j-1].kelamin
			temp3=arData[j-1].golonganDarah
			arData[j-1].umur=arData[j].umur
			arData[j-1].nama=arData[j].nama
			arData[j-1].kelamin=arData[j].kelamin
			arData[j-1].golonganDarah=arData[j].golonganDarah
			arData[j].umur=temp
			arData[j].nama=temp1
			arData[j].kelamin=temp2
			arData[j].golonganDarah=temp3
			if j-2!=-1{
				j=j-1
			}
			
			fmt.Println(j)
			//fmt.Println(temp1," ",temp1," ",temp2," ",temp3)
		}
		
	}

}

func inputData(arData *aData,n *int) {
	fmt.Scan(&arData[*n].umur,&arData[*n].nama,&arData[*n].kelamin)
	fmt.Scanf("%c",&arData[*n].golonganDarah)
	fmt.Println(arData[*n].umur !=-1  ,arData[*n].nama!="kosong",arData[*n].kelamin!="none",arData[*n].golonganDarah!='x',string(arData[*n].golonganDarah))
	for arData[*n].umur !=-1 || arData[*n].nama!="kosong"||arData[*n].kelamin!="none"||arData[*n].golonganDarah!='x' {
		*n=*n+1
		fmt.Scan(&arData[*n].umur,&arData[*n].nama,&arData[*n].kelamin)
		fmt.Scanf("%c",&arData[*n].golonganDarah)
		//fmt.Println(arData[*n].umur !=-1  ,arData[*n].nama!="kosong",arData[*n].kelamin!="none",arData[*n].golonganDarah!='x',string(arData[*n].golonganDarah))
	}

}
func carMaxMinUmur (arData aData,n int) (int,int) {
	var indx,indx2 int
	for i:=1;i<n;i++{
		if arData[indx].umur<arData[i].umur{
			indx=i
		}else if arData[indx2].umur>arData[i].umur {
			indx2=i
		
		}
	}
	return indx,indx2

}

func cariUmur(arData aData,n ,umurDicari int) int {
	var l,r, mid ,index int
	index =-1
	l=0
	r=n-1
	mid = (l+r)/2
	for l<=r && index==-1{
		if umurDicari<arData[mid].umur{
			r=mid-1
		}else if umurDicari>arData[mid].umur{
			l=mid+1
		}else {
			index=mid
		}
		mid = (l+r)/2
	}
	//fmt.Println(index)
	return index
}
/*3 ahmad cowo a 8 Juber cowo b 11 setya cewe o 15 putri cewe c 17 asep cowo a 24 leni cewe a 29 saipudin cowo a*/
/*-1 kosong none x*/

/*8 ahmad cowo a 9 Juber cowo b 13 setya cewe o 12 putri cewe c 11 asep cowo a 16 leni cewe a 20 saipudin cowo a*/
