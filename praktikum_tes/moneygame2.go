package main
import "fmt"
/*program ini menentukan jumlah pecahan 10.000 dalam suatu angka yang dimasukkan dan harus lebih dari 10.000 itu sendiri
karena jika tidak program akan keluar*/
func main(){
	var (a,b int)

	fmt.Scan(&a)
	for a>10000{
		b=a/10000
		fmt.Println("Banyak Uang Rp10.000,- : ",b)
		fmt.Scan(&a)
	}
}