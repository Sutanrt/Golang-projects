package main 
import "fmt"
type x [6]int

func main () {
	var ar1,ar2 x
	var pindah,pindah2, n int
	for i:=0;i<6;i++{
		ar1[i]=i
		
	}
	fmt.Println(ar1)
	n=0
	for i:=0;i<3;i++{
		fmt.Scan(&pindah,&pindah2)
	pindahIsi(&ar1,&ar2,pindah,pindah,&n)
	
	pindahIsi(&ar1,&ar2,pindah2,pindah,&n)
	tambahArr(&ar1,pindah2,&n)
	tampilArr(ar1,len(ar1)-n)
	tampilArr(ar2,6)
	
	
	}
}


func pindahIsi (ar1,ar2 *x,pindah,pindah2 int,n *int) {
	ar2[pindah]=1
	fmt.Println(pindah)
	if pindah2<pindah{
		for i:=pindah-*n;i<len(ar1)-1;i++{
		ar1[i]=ar1[i+1]
	}
	}else{
		for i:=pindah;i<len(ar1)-1;i++{
		ar1[i]=ar1[i+1]
	}
	}
	
	*n=*n+1

}
func tambahArr (ar1 *x,pindah2 int,n *int) {
	ar1[(5-*n)+1]=pindah2
	*n=*n-1


}

func tampilArr(ar1 x, n int) {
	for i:=0;i<n;i++{
		fmt.Print(ar1[i]," ")
	}
	fmt.Println()
}

func cari () {
	


}




















































































































































