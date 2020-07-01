package main
import "fmt"
import "strconv"

func main() {

		var str = "10"
		//str to int
		var strtoint, err = strconv.Atoi(str)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(strtoint) //10 int

		//int to str
		var inttostr = strconv.Itoa(strtoint)
		fmt.Println(inttostr)//10 string
		//float to str
		var flt = 30.7
		var floattostr = strconv.FormatFloat(flt,'f', 6, 64) //30.700000 string
		fmt.Println(floattostr)

		//str to float
		var strtoflt, err2 = strconv.ParseFloat(floattostr, 8) //30.700000762939453 32byte 30.7 8byte

		if err2 == nil {
			fmt.Println(strtoflt)
		}

}
