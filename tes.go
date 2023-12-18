package main 
import "fmt"

func denom(uang int,k10,k2,k3 *int) {
	*k10=uang/10000
	uang=uang%10000
	*k2=uang/2000
	uang=uang%2000
	*k3=uang/1000


} 
func main () {
	var uang, k10,k2,k3 int
	fmt.Scan(&uang)
	denom(uang,&k10,&k2,&k3)
	fmt.Println(k10,k2,k3)
}


/*func cek (a,b int , min,max *int) {
	if a>b{
		*min =b
		*max=a
	}else {
		*min=a
		*max=b
	}
}
func main () {
	var a,b,c,d,max1,min1,max2,min2,temp int
	fmt.Scan(&a,&b,&c,&d)
	cek(a,b,&min1,&max1)
	cek(c,d,&min2,&max2)
	cek(min1,min2,&min1,&temp)
	cek(max1,max2,&temp,&max1)
	fmt.Println(min1,max1)
	
	
}*/


/*func Luas (a,t int,luas *float64) {
	
	*luas = float64(a*t)/2
	reference


}

func main () {
	var luas float64
	var a,t int

	fmt.Scan(&a,&t)
	Luas(a,t,&luas)
	//scan basically return
	fmt.Println(luas)

}*/