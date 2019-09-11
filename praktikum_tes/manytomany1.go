package main
import "fmt"
/*program ini meminta jumlah loop untuk ahmad dan badrun, variabel untuk ahmad berbeda dengan badrun,
p banyak loop,n counter, r rata-rata, t total, a input nilai, b boolean
variabel dengan angka satu milik badrun*/
func main(){
	var (n,p int
		a,r,t float64
		n1,p1 int
		a1,r1,t1 float64
		b bool)
	fmt.Print("Banyak Loop Ahmad  : ")
	fmt.Scanln(&p)
	fmt.Print("Banyak Loop Badrun  : ")
	fmt.Scanln(&p1)
	
	//Ahmad
	n=1
	r=0
	p=p*3
	for n<=p {
		fmt.Scan(&a)
		t=t+a
		n++
		
	}
	r=t/float64(p)
	//Badrun
	n1=1
	r1=0
	p1=p1*3
	for n1<=p1 {
		fmt.Scan(&a1)
		t1=t1+a1
		n1++
	}
	r1=t1/float64(p1)
	b=r>r1
	fmt.Println("Rata-rata Ahmad  ", r, " . Menang? ", b)
	fmt.Println("Rata-rata Badrun ", r1, ". Menang? ", !b)
}