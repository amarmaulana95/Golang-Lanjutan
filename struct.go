package main
import "fmt"

type pelajar struct {
	nama string
	kelas int
	umur int
}

func main() {
	var x1 = pelajar{"Johan", 12, 17}
	fmt.Println(x1.nama)
	fmt.Println(x1.kelas)
	fmt.Println(x1.umur)
}
