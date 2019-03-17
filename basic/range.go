package main
import "fmt"
func main() {
// 	nums:=[]int{2,4}
// 	sum:= 0
// 	for _, num:= range nums{
// 		sum+=num
// 	}
// 	fmt.Println(sum)
product := map[string]string{"a":"apple","b": "banana","c": "cat"}
for item, value := range product{
	fmt.Printf("%s -> %s\n",item, value)
}
}