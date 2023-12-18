package main 

import "fmt"

func faktorial (a *int)  {
	if *a==0 {
		*a=1
	}else {
	for i:=int(*a)-1 ;i>1;i--{
		*a=int(*a)*i
	} 
	}

}
func permutation (a,b int,jumlah *int)  {
	 c :=a-b
	faktorial(&a)
	faktorial(&c)
	*jumlah = a/c
	
}

func combination (a,b int, jumlah *int)  {
	var c int
	c=a-b
	faktorial(&a)
	faktorial(&b)
	faktorial(&c)
	*jumlah=a/(b*c)

}
func main () {
	var a,b,c,d,jumlah,jumlah1 int 
	fmt.Scan(&a,&b,&c,&d)
	permutation(a,c,&jumlah)
	combination(a,c,&jumlah1)
	fmt.Println(jumlah,jumlah1)
	permutation(b,d,&jumlah)
	combination(b,d,&jumlah1)
	fmt.Println(jumlah,jumlah1)
	
	
	
}