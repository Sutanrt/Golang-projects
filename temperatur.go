package main
import (
	"fmt"
)
func main () {
	var celcius,fahrenheit float64
	fmt.Scan(&celcius)
	fahrenheit=celcius*9/5-32
	fmt.Print(fahrenheit)

}