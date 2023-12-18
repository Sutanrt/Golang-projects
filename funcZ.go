package main 
import ("fmt";"math")
func z_1301223465 (x,y int) float64 {
	return (math.Sqrt(float64(3*x*6*y)/float64(4*y)))

}
func main () {
	var x,y int
	fmt.Scan(&x,&y)
	fmt.Printf("%.3f %.3f",z_1301223465(x,y),z_1301223465(y,x))

}