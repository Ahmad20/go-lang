package main
import "fmt"
/*program ini menentukan pemenang dari perbandingan dengan nilai min, p merupakan jumlah loop yang dikali 3
t adalah total nilai, a input nilai, r rata-rata, b boolean*/
func main(){
	var (min,a,r,t float64
		p,n int
		b bool)

	fmt.Scanln(&min)
	fmt.Scan(&p)

	for p!=0{
		n=1
		r=1
		t=0
		p=p*3
		for n<=p{
			fmt.Scan(&a)
			t=t+a
			n++
		}
		r=t/float64(p)
		b=r>min
		fmt.Println("Rata-rata peserta ",r,"Menang ? ",b)
		fmt.Scan(&p)
		
	}
	
	
}