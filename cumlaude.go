package  main 
import "fmt"

func cumlaude_1301223465(ipk float64,masaStudi int,publikasi bool) bool {
	return ipk>=3.51&&ipk<=4.00 && masaStudi<=8&&publikasi
}
func sangatMemuaskan_1301223465(ipk float64,masaStudi int,publikasi bool)bool{
	return ipk>=2.76&&masaStudi<=14
}
func memuaskan_1301223465 (ipk float64,masaStudi int, publikasi bool)bool{
	return ipk>=2.00&&masaStudi<=14
}
func main ()  {
	var ipk float64
	var masaStudi int
	var publikasi bool
	fmt.Scan(&ipk,&masaStudi,&publikasi)
	if cumlaude_1301223465(ipk,masaStudi,publikasi){
		fmt.Print("Cumlaude")
	}else if sangatMemuaskan_1301223465(ipk,masaStudi,publikasi){
		fmt.Print("Sangat memuaskan")
	}else if memuaskan_1301223465(ipk,masaStudi,publikasi){
		fmt.Print("Memuaskan")
	}
}