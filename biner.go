package main 
import "fmt"

func main () {
	var x int
	var angkaS string
	fmt.Scan(&x)
	for x>0 {
	
		angkaS=angkaS+fmt.Sprint(x%2)
		x=x/2
	}
	for i:=len(angkaS)-1;i>=0;i--{
		fmt.Print(string(angkaS[i]))
	}
}