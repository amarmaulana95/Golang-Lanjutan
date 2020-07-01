package main
import "fmt"

func main() {
  // array
  var x [5]float64
  x[4] = 100
  fmt.Println(x) //[0 0 0 0 100]
}
