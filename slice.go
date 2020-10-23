package main

import "fmt"

func main() {
	var things = []string{"meleto", "ropa", "roloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)
	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "superman"
	fmt.Println(heros)
	heros = append(heros, "deadpool")
	fmt.Println(heros)
	fmt.Println(cap(heros))

}
