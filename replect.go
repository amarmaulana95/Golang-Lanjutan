package main
import "fmt"
import "reflect"

func main() {
	var number = "10"
	var reflectnumber = reflect.ValueOf(number)

	fmt.Println("type variabel: ", reflectnumber.Type())

	if reflectnumber.Kind() == reflect.Int{ //jika typedata integer
		fmt.Println("Nilai variabel: ", reflectnumber.Int())
	}
}
