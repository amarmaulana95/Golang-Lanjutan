package main
import "fmt"
import "encoding/base64"

func main() {
	var nama_saya = "nida shofa fauziah"
	var encodestring = base64.StdEncoding.EncodeToString([]byte(nama_saya))
	fmt.Println("encoding string ", encodestring) // merahasiakan nama menjadi suatu test yang acak, diubah format nya jadi kode encripsi base64
	//decoding
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodestring)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
