package main 
import "fmt"

const NMAX int =1000

type himpunan struct {
	info [NMAX] string
	nElemen int
}

func main () {
	var A,B,C,D himpunan
	
	createSet_1301223465(&A)
	createSet_1301223465(&B)
	intersection_1301223465(A,B,&C)
	union_1301223465(A,B,&D)
	
	printSet_1301223465(C)
	fmt.Println()
	printSet_1301223465(D)
}
func createSet_1301223465 (set *himpunan) {
	var s string
	fmt.Scan(&s)
	for !isMember_1301223465(*set,s) {
		set.info[set.nElemen]=s
		set.nElemen=set.nElemen+1
		fmt.Scan(&s)
	}	
}
func isMember_1301223465(set himpunan, s string) bool {
	var cek bool
	var i int
	for !cek && i<=set.nElemen{
		cek=s==set.info[i]
		i=i+1
	} 
	return cek
}
func intersection_1301223465(set1,set2 himpunan, set3 *himpunan) {
	var idx,idx3 int

	for idx<=set1.nElemen {
		if isMember_1301223465(set2,set1.info[idx]){
			set3.info[idx3]=set1.info[idx]
			idx3=idx3+1
		}
		idx=idx+1
	}
	set3.nElemen=idx3-1
}
func union_1301223465(set1,set2 himpunan, set3 *himpunan) {
	var idx int
	
	for idx<=set1.nElemen {
		set3.info[idx] =set1.info[idx]
		idx=idx+1
	}
	
	set3.nElemen=idx-1
	idx=0
	for idx<=set2.nElemen {
		if !isMember_1301223465(*set3,set2.info[idx]){
			set3.info[set3.nElemen]=set2.info[idx]
			set3.nElemen=set3.nElemen+1
		}
	idx=idx+1
	}		
	
}

func printSet_1301223465(set himpunan) {
	for i:=0;i<set.nElemen;i++{
		fmt.Print(set.info[i]," ")
	}
}

