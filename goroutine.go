package main
import "fmt"
import "runtime"

func tampilkan_pesan(x int, pesan string){
	for i := 0; i < x; i++ {
		fmt.Println((i +1), pesan)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go tampilkan_pesan(5, "saya kirim")
	tampilkan_pesan(5, "saya kedua")

	var input string	//untuk bloking, sehingga kedua nya dijalankan
	fmt.Scanln(&input)
}
