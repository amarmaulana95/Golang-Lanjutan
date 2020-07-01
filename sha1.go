package main
import "crypto/sha1"
import "fmt"

func main() {
	var text = "text ini rahasia"
	var sha = sha1.New() // object bertype hash  has
	sha.Write([]byte(text))
	var encripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", encripsi)

	fmt.Println(stringenkripsi) //e83a5494f88a7fdd0d92b4d110af1e0552ae4380
}
