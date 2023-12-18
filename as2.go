package main 
import "fmt"
func valid_1301223465 (id int) bool {
	var i int
	for id >0 {
		i=i+1
		id = id/10
	}
	return i==10
}
func respond_1301223465 (id int, n string, g rune) {
	if g=='f'{
		fmt.Print("Mrs. ")
	}else if g=='m'{
		fmt.Print("Mr. ")
	}
	fmt.Print(n," ID:",id)
	if valid_1301223465(id){
		fmt.Print (" anda valid")
	}else {
		fmt.Print (" anda tidak valid")
	}

}
func main () {
	var id int
	var n string
	var g rune
	fmt.Scanf("%d %s %c",&id,&n,&g)
	respond_1301223465(id,n,g)
}
