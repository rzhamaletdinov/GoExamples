package main

import "fmt"
import "mymodule"

func main() {
	xs := []float64{1,2,3,4}
	avg := mymodule.Average(xs)
	fmt.Println(avg)
}