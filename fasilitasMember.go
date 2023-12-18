package main 
import "fmt"

func main () {
	var total_belanja int
	var status_membership string
	fmt.Scan(&total_belanja,&status_membership)
	if status_membership=="Yes" {
		if total_belanja>100000{
			fmt.Print("Anda mendapat discount 5%, tambahan poin 10, dan free gift ")
		}else if total_belanja>75000{
			fmt.Print("Anda mendapat discount 5% dan tambahan poin 5")
			
		}else{
			fmt.Print("Anda mendapat tambahan poin 5")
		}
	}

}