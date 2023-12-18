package main
import "fmt" 
func jam (a int) int {
	return a*3600
}
func menit (a int)int{
	return a*60
}

func main () {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	fmt.Print(menit(b)+jam(a)+c)

}