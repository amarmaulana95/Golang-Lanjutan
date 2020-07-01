package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	var number = []int{1,2,3,4,5,67,5}
	fmt.Println("Number ", number)

	var channel1 = make(chan float64) 	//u/ menghitung nil rata2
	go getratarata(number, channel1) //kirim channel1
	var channel2 = make(chan int)		//u/ hitung nil max
	go nilaimaksimal(number, channel2)

	for i := 0; i < 2; i++ {
		select {
			case rata_rata := <- channel1:
			fmt.Println("Rata-rata \t : %.2f \n", rata_rata)
			case maksimal := <- channel2:
			fmt.Println("Maksimal \t : %d \n", maksimal)
		}
	}
}

func getratarata(number []int, ch chan float64){
	var sum = 0
	for _, e := range number{
		sum += e
	}

	ch <- float64(sum) / float64(len(number))
}

func nilaimaksimal(number []int, ch chan int){
	var max = number[0]
	for _, e := range number{
		if max < e{
			max = e
		}
	}
	ch <- max
}
