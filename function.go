package main
import "fmt"
// contoh function
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
	 total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
		total := 0
		for _, v := range args {
		  total += v
		}
		return total
}

func factorial(x  uint) uint {
		if x == 0 {
		 return 1
		}
		return x * factorial(x-1) //2*(2-1)
}

func main() {
	xs := []float64{1,2,3,4}
  // function
	fmt.Println(average(xs)) //2.5
  // variadic
  fmt.Println(add(1,2,3)) // 6
  rs := []int{1,2,3}
	fmt.Println(add(rs...)) // using slice, return 6
  // closure
  add := func(x, y int) int {
			return x + y
	}
	fmt.Println(add(1, 1)) //2

	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment ()) //1
	fmt.Println(increment ()) //2
  // recursion
  fmt.Println(factorial(2))//2
  fmt.Println(factorial(0))//1
  fmt.Println(factorial(3))//6
}
