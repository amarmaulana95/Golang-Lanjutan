package math
	// find the average of a series of numbers
	func Average(xs []float64) float64 {
		total := float64(0)
		for _, x := range xs {
			total += x
		}
		return total / float64(len(xs))
	}
// go install - to compile math.go
// godoc 
