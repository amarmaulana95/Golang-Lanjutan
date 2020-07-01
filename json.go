package main
	import "encoding/json"
	import "fmt"

	type User struct{
		Nama string "json:Nama"
		Umur int
	}

	func main() {
		/*
		var jsonString = `{"Nama" : "ilham", "Umur" : 12}`

		var jsonData = []byte(jsonString)

		var data User

		var err = json.Unmarshal(jsonData, &data) //menghasilkan json object
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Nama ", data.Nama)
		fmt.Println("Umur ", data.Umur)
		*/
		//json string
		var object = []User{{"Ilham", 13},{"Abdul", 26}}
		var jsonData2, err = json.Marshal(object)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var jsonString2 = string(jsonData2)
		fmt.Println(jsonString2) //[{"Nama" : "ilham", "Umur" : 13}, {...}]
	}
