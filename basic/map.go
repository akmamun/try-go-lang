package main
import "fmt"
func main() {
	m:=make(map[string]int)
	m["k1"] =1
	m["k2"]=3
	fmt.Println(m)
	fmt.Println("len:", len(m))
	delete(m, "k2")
	fmt.Println("map:", m)
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}