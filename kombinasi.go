package main 

import "fmt"

func faktorial (a int) int {
	if a==0 {
		return 1
	}else {
	for i:=a-1 ;i>1;i--{
		a=a*i
	} 
	}
	return a
}
func permutation (a,b int) int {
	return faktorial(a)/faktorial(a-b)
}

func combination (a,b int) int {
	return faktorial(a)/(faktorial(b)*faktorial(a-b))

}
func main () {
	var a,b,c,d int 
	fmt.Scan(&a,&b,&c,&d)
	fmt.Printf("%d %d\n%d %d",permutation(a,c),combination(a,c),permutation(b,d),combination(b,d))
	
}