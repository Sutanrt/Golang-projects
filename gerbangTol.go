package main
import "fmt"

func main () {
	var x string
	var a,b,c int
	fmt.Scan(&x)
	for x=="A"||x=="B"||x=="C"{
		
		if x=="A" {
			a=a+1
		}else if x=="B" {
			b=b+1
		}else if x=="C" {
			c=c+1
		}
		fmt.Scan(&x)
	}
	fmt.Printf("Tipe A = %v\nTipe B = %v\nTipe C = %v",a,b,c)
}