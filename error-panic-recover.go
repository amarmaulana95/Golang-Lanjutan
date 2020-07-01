package main
	import "fmt"
	import "strconv"

	func main(){
		defer pulihkan_saya()
		var input string
		fmt.Print("Masuskan angka: ")
		fmt.Scanln(&input)

		var number int
		var err error

		number, err = strconv.Atoi(input)
		if err == nil {
			fmt.Println(number, "ini adalah angka")
		}else {
			fmt.Println(input, "bukan angka")
			// fmt.Println(err.Error())
			panic(err.Error())
			fmt.Println("munculkan saya")
		}
	}

	func pulihkan_saya(){
		if r := recover(); r != nil{
			fmt.Println("akhirnya saya ditampilkan")
		}else{
			fmt.Println("ini lancar saja")
		}
	}
