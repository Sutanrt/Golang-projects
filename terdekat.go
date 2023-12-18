package main 

import ("fmt";"math")
type point struct{
	x float64
	y float64
}
func main () {
	var a ,b,c,d point
	fmt.Scan(&a.x,&a.y,&b.x,&b.y,&c.x,&c.y,&d.x,&d.y)
	if jarak_1301223465(a,b)<jarak_1301223465(c,d){
		fmt.Printf("Titik terdekat adalah titik A(%1.f,%1.f) dengan B(%1.f,%1.f) dengan jarak %.1f.",a.x,a.y,b.x,b.y,jarak_1301223465(a,b))
	}else {
		fmt.Printf("Titik terdekat adalah titik C(%1.f,%1.f) dengan D(%1.f,%1.f) dengan jarak %.1f.",c.x,c.y,d.x,d.y,jarak_1301223465(c,d))
	}
}
func jarak_1301223465 (titik,titik2 point) float64 {  
	return math.Sqrt(((titik.x-titik2.x)*(titik.x-titik2.x))+((titik.y-titik2.y)*(titik.y-titik2.y)))
	
}
