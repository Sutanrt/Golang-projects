package main 
import "fmt"

func len_1301223465(num int) int {
	var i int
	for num >0 {
		i=i+1
		num=num/10
	}
	
	return i
}

func pangkat_1301223465 (n int ) int {
	b:=1
	for i:=0;i<n;i++{
		b=b*10
	}
	return b
}

func split_1301223465(num int, num1,num2 *int) {
	var n int
	n=len_1301223465(num)
	
	*num1=num/pangkat_1301223465(n/2)
	*num2=num%pangkat_1301223465(n/2)

}
func main () {
	var num1,num2, num int 
	fmt.Scan(&num)
	split_1301223465(num,&num1,&num2)
	fmt.Println(num1," ",num2)
	fmt.Print(num1+num2)
}