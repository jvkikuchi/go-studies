package main
import "fmt"

func Hello(name string) string  {
	const greeting = "Hello, "

	if (name == ""){
		return greeting + "World"
	}
	
	return greeting + name
}

func main() {
	fmt.Println(Hello("Testing"))
}