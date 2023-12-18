package main 
import "fmt"
func proc (soal, skor *int){
	var s int
	for i:=1;i<=8;i++{
		fmt.Scan(&s)
		
		if s<301 {
			*soal=*soal+s
			*skor=*skor+1
			
		}
	}
}

func main () {
	var s,p string 
	var sl,tw,max1,max2 int
	max1=0

	fmt.Scan(&s)
	for s!="Selesai"{
		if s!="Selesai"{
			sl=0
			tw=0
			proc(&sl,&tw)
			if max1<=tw{
				max1=tw
				max2=sl
				p=s
			}else if max1==tw{
				if max2<sl {
					max2=sl
					p=s
				}
			}
			
		}
		fmt.Scan(&s)
	}
	fmt.Print(p," ",max1,max2)
	
}