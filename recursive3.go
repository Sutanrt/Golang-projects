package main 
import "fmt"

func recursive (s string,b bool, n int) bool{
	if n==len(s)-1-n ||(len(s)-1-n)/2==1{
		//fmt.Println(string(s[n]),string(s[len(s)-1-n]),n,len(s)-1-n,b)
		return b&&string(s[n])==string(s[len(s)-1-n])
	}
	//fmt.Println(n==len(s)-1-n ||n -len(s)==1)
	//fmt.Println(n,len(s)-1-n,n-len(s)-1==1,n-len(s)-1)
	//fmt.Println(string(s[n]),string(s[len(s)-1-n]),n,len(s)-1-n,b)
	return recursive(s,b&&string(s[n])==string(s[len(s)-1-n]),n-1)
}
func main () {
	var s string
	fmt.Scan(&s)
	b:= true
	b=recursive(s,b,len(s)-1)
	if b {
		fmt.Print("YA")
	}else {
		fmt.Print("BUKAN")
	}
}