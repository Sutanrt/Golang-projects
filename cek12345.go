package main 
import "fmt"
type nasi struct {
	a,b,c,d int

}

type nasi2 [3]nasi 

func main () {
	var nasi3,nasi1,temp nasi2
	fmt.Scan(&nasi3[0].a,&nasi3[0].b,&nasi3[0].c,&nasi3[0].d)
	fmt.Scan(&nasi1[0].a,&nasi1[0].b,&nasi1[0].c,&nasi1[0].d)
	fmt.Println(nasi3)
	fmt.Println(nasi1)
	temp=nasi3
	nasi3[0]=nasi1[0]
	nasi1[0]=temp[0]
	fmt.Println(nasi3)
	fmt.Println(nasi1)

}