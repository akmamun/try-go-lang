package main

import "fmt"

func values() (int, int) {
	return 7, 9 // return multi value
}
func main() {
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

}
