package main 
import "fmt"
func main () {
var a,b,i,c,temp int
	i=1
	
	fmt.Scan(&a,&b,&c)
	temp=a
	fmt.Println(46%497)
	for (a-b)%c!=0 {
		fmt.Println((a%c))
		a=temp+a
		i=i+1
		fmt.Println(i,a)
	} 
	fmt.Println(i,a)
	
}