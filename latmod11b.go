package main
import "fmt"
func main () {
	var x [21]int
	var max,n,sah,tidakSah, max2 int
	fmt.Scan(&n)
	for n!=0 {
		if n>0 && n<=20 {
			sah=sah+1
			x[n]=x[n]+1
		}else {
			tidakSah=tidakSah+1
		}
		fmt.Scan(&n)
	}
	
	max=x[0]
	fmt.Println(x)
	for i:=1;i<=20;i++{
		if x[i]>0{
			if x[max] < x[i] ||x[max2]<x[i]{
				fmt.Println("Y ",i)
				if x[max] <x[i]{
					max=i	
				}else {
					max2=i
				}
			}
		}
	
	}
	fmt.Println("Ketua RT: ",max)
	fmt.Println("Wakil ketua: ",max2)
	
}