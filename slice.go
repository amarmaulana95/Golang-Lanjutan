package main
import (
  "fmt"
)

func main() {
  // slice
  slice1 := []int{1,2,3,4,5,6,7,8,9,0}
  x := make([]int, 5, 10)
  fmt.Println(x) //[0 0 0 0 0]
  arr := []int{1,2,3,4,5}
  x = arr[:]
  fmt.Println(x) //[1 2 3 4 5]
  x = arr[0:3]
  fmt.Println(x) //[1 2 3]
  x = arr[3:len(arr)]
  fmt.Println(x) //[4 5]
  copy(slice1, x)
  fmt.Println(x, slice1) //[4 5]
}
