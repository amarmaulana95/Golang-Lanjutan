package main
import "fmt"
//struk
type pelajar struct {
	nama string
	umur int
}
// method
func (s pelajar) namasaya() { //variabel s akan menampung nilai yang telah dideklarasikan
	fmt.Println("nama saya adalah: ", s.nama)
	fmt.Println("umur: ", s.umur)
}

func main() {
	var s1 pelajar //var s1 untuk struct pelajar
	s1.nama = "aldo"
	s1.umur = 19
	var s2 = pelajar{"yosep", 18} //deklarasi diawal

	s1.namasaya()
	s2.namasaya()
	// pengambilan data dienkapsulasi sehingga lebih mudah
}
/*
func (c *Circle) area() flaot64 { //  (c *Circle) receiver: like a parameter it has a name and a type
	return math.Pi * c.r*c.r
}
*/
