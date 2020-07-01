package main
import "fmt"

func main() {
var nomor int = 10
var alamat_nomor *int = &nomor

fmt.Println("Nilai dari nomor: ", nomor)
fmt.Println("alamat variabel nomor: ", alamat_nomor)
}
