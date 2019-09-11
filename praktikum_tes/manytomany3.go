package main
import "fmt"
/*program tantangan ini memerintahkan user memasukkan baris yang ditujukan pada jumlah peserta,
p jumah nilai yang ingin dimasukkan, min nilai terendah, peserta pemenang merupan yang reratanya melewati min(b) */
func main(){
	var (bar,p,i,j int
		min,a,r,t float64
		b bool)

	fmt.Scan(&bar)
	fmt.Scan(&p)
	fmt.Scan(&min)
	i=1
	for i<=bar {
		t=0
		j=1
		r=0
		p=p*3
		fmt.Print("Nilai ",i,":")
		for j<=p {
			
			fmt.Scan(&a)
			t=t+a
			r=t/float64(p)
			b=r>min
			
			j++

		}
		p=p/3
		fmt.Println("Peserta,",i," Menang ? ",b)
		fmt.Println("")
		i++
		
	}
	
	
}