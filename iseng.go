package main 
import "fmt"
const NMAX int = 32
type data1 struct {
	tgl int
	hrg float64
	
}
type data [NMAX]data1
func main () {
	var c data
	var n,tglTarget int
	var index float64
	index=-1
	//isiArray(&c,&n)
	isiArray2(&c,&n)
	fmt.Scan(&tglTarget)
	//fmt.Println(cari(c,32,tglTarget))
	//fmt.Println(cari2(c,n,tglTarget))
	fmt.Println(c,n)
	fmt.Println(carRecursive(c,tglTarget,16,0,n,index))
}
func isiArray (c *data,n *int) {
	var a int
	var b float64
	fmt.Scan(&a,&b)
	for a!=0 || b!=-1.0 {
		c[a].hrg=b
		c[a].tgl=a
		*n=*n+1
		fmt.Scan(&a,&b)
	}
	//fmt.Println(c)

}
func cari (c data,j,n int) float64 {
	var l,m,r,index int
	l=0
	r=j-1
	m=(l+r)/2
	index=-1
	for index==-1 && l<=r {
		//fmt.Println("HH",l," ",m," ",r," ",n,c[m].tgl)
		if n <m{
			r=m-1
		}else if n>m{
			l=m+1
		}else{
			index=m
		}
		m=(l+r)/2
	}
	return c[index].hrg
}
func isiArray2 (c *data,n *int) {
	var a int
	var b float64
	fmt.Scan(&a,&b)
	for a!=0 || b!=-1.0 {
		c[*n].hrg=b
		c[*n].tgl=a
		*n=*n+1
		fmt.Scan(&a,&b)
	}
	//fmt.Println(c)

}
func cari2 (c data,j,n int) float64 {
	var l,m,r int
	var index float64
	l=0
	r=j-1
	m=(l+r)/2
	index=-1
	fmt.Println(c)
	fmt.Println(n)
	for index==-1 && l<=r {
		fmt.Println(l,m,r)
		if n <c[m].tgl{
			r=m-1
		}else if n>c[m].tgl{
			l=m+1
		}else{
			index=c[m].hrg
		}
		m=(l+r)/2
	}
	return index
}
func carRecursive (c data,n,m,l,r int,index float64) float64{
	if l>r || index!=-1{
		return index
	}else {
		if n<c[m].tgl {
			m=(l+r)/2
			return carRecursive(c,n,m,l,m-1,index)
		}else if n>c[m].tgl{
			fmt.Println("H1")
			m=(l+r)/2
			return carRecursive(c,n,m,m+1,r,index)
		}else {
			index=c[m].hrg
			fmt.Println("H")
			return carRecursive(c,n,m,l,r,index)
		}
		
	}
}
