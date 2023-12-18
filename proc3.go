package main 
import "fmt"

func proc (n int,h *int){
	
	if n==0 {
		*h=1
	}else {
		for i:=n-1;i>=1;i--{
			n=n*i
		}
		*h=n
	}
	
	

}

func perm(n, r int ,h *int) {
	var c,c1 int

	proc(n,&c1)
	proc(n-r,&c)
	//fmt.Println(c1,c)
	*h=c1/c
	
}

func comb(n, r int ,h *int) {
	var c,c1,c2 int

	proc(n,&c1)
	proc(n-r,&c2)
	proc(r,&c)
	//fmt.Println(c,c1,c2,"halohai")
	*h=c1/(c2*c)
	
}
func main () {
	var a,b,c,d, h,h1 int
	fmt.Scan(&a,&b,&c,&d)
	perm(a,c,&h)
	comb(a,c,&h1)
	fmt.Println(h,h1)
	perm(b,d,&h)
	comb(b,d,&h1)
	fmt.Print(h,h1)
	
}