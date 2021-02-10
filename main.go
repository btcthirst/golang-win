package main

import "fmt"

func main() {
	var hello string
	a, b, c := 5, 2.5, "some"

	_, err := fmt.Scanln(&hello)
	if err != nil {
		panic(err)
	}

	fmt.Println(hello)
	fmt.Println(a, b, c)
	fmt.Println(a + int(b))
	fmt.Println(b + float64(a))
	fmt.Println("hello World! and ", hello)
	fmt.Println("hello World!")
}
