package main 
import "fmt"
func main () {
var x [10]int
for j:=0;j<10;j++{
	fmt.Scan(&x[j])
}
fmt.Println(sortingDesc(x))
fmt.Println(sortingAsc(x))
fmt.Println(insSelectAsc(x))
fmt.Println(insSelectDSC(x))
}

func sortingAsc (x [10]int) [10]int {
	var i,idx,c int
	for i<10 {
		idx=i
		for j:=i;j<10;j++{
			if x[idx]>x[j] {
				idx=j
			}
			
		}
		c=x[i]
		x[i]=x[idx]
		x[idx]=c
		i++
		//fmt.Println(idx)
	}
	return x
}
func sortingDesc (x [10]int) [10]int {
	var i,idx,c int
	for i<10 {
		idx=i
		for j:=i;j<10;j++{
			if x[idx]<x[j] {
				idx=j
			}
			
		}
		c=x[i]
		x[i]=x[idx]
		x[idx]=c
		i++
		//fmt.Println(idx)
	}
	return x
}

func insSelectAsc (x[10]int) [10]int {
	var j,temp int
	for i:=1;i<10;j++{
		j=i
		for j>0 && x[j-1]>x[j]{
			temp=x[j-1]
			x[j-1]=x[j]
			x[j]=temp
			j=j-1
		
		}
		i=i+1
	}
	return x
}
func insSelectDSC (x[10]int) [10]int {
	var j,temp int
	for i:=1;i<10;j++{
		j=i
		for j>0 && x[j-1]<x[j]{
			temp=x[j-1]
			x[j-1]=x[j]
			x[j]=temp
			j=j-1
		
		}
		i=i+1
	}
	return x
}