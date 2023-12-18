package main 
import "fmt"
func main () {
	var kartu [55]int
	var i int
	fmt.Scan(&kartu[i])
	for kartu[i]!=0{
		i=i+1
		fmt.Scan(&kartu[i])
		
	}
	i=i-1
	for i>=0{
		fmt.Print(kartu[i]," ")
		i=i-1
	}
}