package main
import "fmt"
import "os"

func main(){
	defer fmt.Println("hai!")
	fmt.Println("selamat datang")
	os.Exit(1) // defer dan tampilkan tidak dijalankan
	tampilkan()
}

func tampilkan(){
	fmt.Println("saya ditampikan")
}
