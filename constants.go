package main

import "fmt"
import "math"
func main() {
	const n = 6000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) 
	// math calculation
	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))

}