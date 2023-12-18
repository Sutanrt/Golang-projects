package main 

import ("fmt";"math")

func t (xi,xj,d int)int{
	var y int
	y=xj-xi
	if y<0{
		y=-y
	}
	return int(math.Pow(float64(y),float64(d)))
	
}

func main () {
	var  a,b,c,d,e int
	min :=10000000000
	max:=-10000000000
	var slc[]int
	var sl[]int
	
	fmt.Scan(&c,&d)
	
	for i:=0;i<c;i++{
		
		fmt.Scan(&a,&b)
		slc=append(slc,a)
		sl=append(sl,b)
	
	}
	for j:=0;j<len(slc);j++{
		for k:=j+1;k<len(slc);k++{
			e=t(slc[j],slc[k],d)+t(sl[j],sl[k],d)
			if min > e {
				min = e
			}
			if max<e {
				max=e
			}
		}
	}
	fmt.Println(min,max)
}