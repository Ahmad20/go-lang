package main
import "fmt"
/*program ini untuk menentukan uang yang valid di indonesia, sebab nilai terkecil di indonesia ialaha 25 perak
maka nilai yang tidak habis dibagi 25 artinya tidak valid, dan program akan berhenti ketika input yang dimasukkan
valid atau dapat dibagi 25*/
func main(){
	var (a int)

	fmt.Scan(&a)
	for a%25 !=0 {
		fmt.Println(a, " Tidak Valid")
		fmt.Scan(&a)
	}
	fmt.Print(a, " Valid")
}