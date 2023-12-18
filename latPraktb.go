package main 
import "fmt"

type set [2022]int 
func exist(T set, n int, val int) bool {
	var c bool
	var i int
	c=true
	for i<n && c{ 
		c=T[i]!=val
		i=i+1
	}
	return c

}
func inputSet(T *set, n *int) {
	var val int
	fmt.Scan(&val)
	for exist(*T , *n, val) && *n<2022{
		T[*n]=val
		*n=*n+1
		fmt.Scan(&val)
	}
}
func findIntersection(T1,T2 set,n,m int,T *set, h *int) {
	for i:=0;i<n;i++{
		if !exist(T2,m,T1[i]){
			T[*h]=T1[i]
			*h=*h+1
		}
	}

}
func PrintSet(T set, n int) {
	for i:=0;i<n;i++{
		fmt.Print(T[i]," ")
	}

}
func main () {
	var T,T1,T2 set
	var n,m,h int
	inputSet(&T1,&n)
	inputSet(&T2,&m)
	findIntersection(T1,T2,n,m,&T,&h)
	PrintSet(T,h)
}