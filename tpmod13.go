package main 
import "fmt"

type student struct {
	name string
	math,indo,eng,sains,average float64
}

type arrStudent [2048]student
func main () {
	var T arrStudent
	var n int
	entryStudent_1301223465(&T,&n)
	ranking_1301223465(&T,n)
	
	printTop3_1301223465(T,n)
	printMax_1301223465(T,n)
}

func entryStudent_1301223465(T *arrStudent, n *int) {
	var nama string
	fmt.Scan(&nama)
	for nama!="STOP" {
		T[*n].name=nama
		fmt.Scan(&T[*n].math,&T[*n].indo,&T[*n].eng,&T[*n].sains)
		calculateAverage_1301223465(T,*n)
		*n=*n+1
		fmt.Scan(&nama)
	}

}

func calculateAverage_1301223465(T *arrStudent,n int) {
	T[n].average=(T[n].math+T[n].indo+T[n].eng+T[n].sains)/4
}

func max_1301223465(T arrStudent,n int,flag string) int {
	var index int
	if flag=="ind"{
		for i:=0;i<n;i++{
		if T[index].indo<T[i].indo{
			index=i
		}
		
	}
	}else if flag=="math"{
		for i:=0;i<n;i++{
			if T[index].math<T[i].math{
				index=i
			}
		
	}
	}else if flag=="eng"{
		for i:=0;i<n;i++{
			if T[index].eng<T[i].eng{
			index=i
		}
		
	}
	}else if flag=="sains"{
		for i:=0;i<n;i++{
			if T[index].sains<T[i].sains{
			index=i
		}
		
	}
	}
	

	return index
}

func ranking_1301223465 (T *arrStudent, n int){
	var temp float64
	var max int
	var temp2 string
	
	for i:=0;i<n;i++{
		max=i
		for j:=i+1;j<n;j++{
			if T[max].average<T[j].average{
				max=j
			}
		}
		temp=T[i].average
		T[i].average=T[max].average
		T[max].average=temp
		
		temp=T[i].math
		T[i].math=T[max].math
		T[max].math=temp
		
		temp=T[i].indo
		T[i].indo=T[max].indo
		T[max].indo=temp
		
		temp=T[i].sains
		T[i].sains=T[max].sains
		T[max].sains=temp
		
		temp=T[i].eng
		T[i].eng=T[max].eng
		T[max].eng=temp
		
		temp2=T[i].name
		T[i].name=T[max].name
		T[max].name=temp2
	}
}

func printTop3_1301223465(T arrStudent, n int) {
	for i:=0;i<3;i++{
		fmt.Printf("Rangking %d: %s rata-rata %.1f\n",i+1,T[i].name,T[i].average)
	}
}

func printMax_1301223465(T arrStudent,n int) {
	fmt.Println("Nilai Matematika tertinggi diraih oleh ",T[max_1301223465(T,n,"math")].name)
	fmt.Println("Nilai Bahasa Indonesia tertinggi diraih oleh ",T[max_1301223465(T,n,"ind")].name)
	fmt.Println("Nilai Bahasa Inggris	tertinggi diraih oleh ",T[max_1301223465(T,n,"eng")].name)
	fmt.Print("Nilai IPA/IPS diraih oleh ",T[max_1301223465(T,n,"sains")].name)
}
/*
Ana 80 90 70 50
Dini 10 30 90 50
Joni 20 80 75 77
Dika 75 71 73 60
Desi 88 45 76 47
Vito 67 78 67 56
Vita 50 86 84 88
Keysa 77 89 61 87
Enda 87 65 77 71
STOP
*/