package main
import "fmt"
func main() {
	s := make([]string, 1)
	fmt.Println("emp:", s)
	s = append(s, "d")
	fmt.Println("item:", s,"len:" ,len(s))
	c := make([]string,2)
	copy(c, s)
	fmt.Println("copy:", c)
}