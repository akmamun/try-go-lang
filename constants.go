package main

import "fmt"
import "math" //for calculation
// func main() // its throw an error can't use new line before '{ ' clause
// {
	// num:=2
	// fmt.Println(num)
// }

func main() {
	// m:=2 //assign value of variable it 
		//throw an error bcz in go can't declard if not use varaiable
		
	const n = 6000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) 
	// math calculation
	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))

}