package repeat

import "fmt"

func Repeat(value string, n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += value
	}

	return result
}

func main(){
	result := Repeat("a", 2)
	fmt.Printf(result)

}