package main

import "fmt"

func main() {
	start := 1
	things := []string{"matter", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])

	}

}
