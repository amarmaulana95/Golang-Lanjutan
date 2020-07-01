package main
import "fmt"
import "strings"

func main(){
	var apakahada = strings.Contains("disini adalah rumah", "ini") //true
  fmt.Println(apakahada)
  fmt.Println(strings.Contains("test", "es")) //true
	fmt.Println(strings.Count("test", "t"))	//2
	fmt.Println(strings.HasPrefix("test", "te"))//true
	fmt.Println(strings.HasSuffix("test","st"))	//true
	fmt.Println(strings.Index("test","e"))	//1
	fmt.Println(strings.Join([]string{"a", "b"},"-"))	//"a-b"
	fmt.Println(strings.Repeat("a",5))	//== "aaaaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))	//"bbaa"
	// fmt.Println(strings.Splite("a-b-c-d-e", "-"))	//[]string{"a","b","c","d","e"}
	fmt.Println(strings.ToLower("TEST"))		// "test"
	fmt.Println(strings.ToUpper("test"))		// "TEST"
}
