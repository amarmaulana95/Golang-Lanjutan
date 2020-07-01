package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	pesan := make(chan int, 3) // jumlah buffer  nya 2
	//mengirim
	go func(){
		for {
			i := <-pesan
			fmt.Println("terima data", i)
		}
	}()
	//menerima
	for i := 0; i < 5; i++ {
		fmt.Println("kirim data ", i) // kirim data ke variabel channel pesan
		pesan <- i
	}
}
