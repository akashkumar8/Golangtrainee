package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {

	rob := User{"rob", "rob@gmail.com", 50}
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%v\n", rob)
	fmt.Printf("%v\n", rob.Email)

	var sam = new(User)
	sam.Email = "sam@gmail.com"
	sam.age = 23
	fmt.Printf("%+v\n", sam)

}
