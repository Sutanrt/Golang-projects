package main 
import "fmt"

const nMax int = 51
type mahasiswa struct {
  NIM int
  nama string
  nilai int 

}

type aMhs[nMax]mahasiswa

func main () {
	var mhs aMhs 
	var n,nim int
	isiData(&mhs,&n)
	fmt.Println("NIM")
	fmt.Scan(&nim)
	fmt.Println(mhs[cariData(mhs,nim,n)].nilai)
	fmt.Println(mhs[cariMaxData(mhs,n,nim)].nilai)
	
}

func isiData (mhs *aMhs, n *int) {
	var nim int
	fmt.Scan(&nim)
	for nim!=-1{
		mhs[*n].NIM=nim
		fmt.Scan(&mhs[*n].nama,&mhs[*n].nilai)
		*n=*n+1
		fmt.Scan(&nim)
	}
}


func cariDataSequental (mhs aMhs, nim,n int) int {
	var i int
	var idx=-1
	for i<n && idx==-1 {
		if mhs[i].NIM==nim{
			idx=i
		}else {
			i=i+1
			
		}
	
	}
	return idx
	
	
}

func cariDataBinary(mhs aMhs,nim,n int) {
	var l,r,mid,idx int
	idx=-1
	l=0
	r=n-1
	mid=(l+r)/2
	for l<=r && idx==-1{
		
	
	}
	
	

}
func cariData(mhs aMhs, nim,n int) int {

	var i int
	var idx=-1
	
	for i<n&&idx==-1 {
		if mhs[i].NIM!=nim {
			i=i+1
		}else {
			idx=i
		}
		
	}
	return idx
}

func cariMaxData (mhs aMhs, n, nim int) int {
	var i,idx int
	idx=cariData(mhs,nim,n)
	for i < n{
		if mhs[i].NIM==nim{
			if mhs[idx].nilai<mhs[i].nilai{
				idx=i
			}
		}
		i=i+1
	}
	return idx
}
