package main 

import "fmt"

type User struct {
        Name string
        Email string
        Age int
        Roll number int

    }
    func main() {
        ak := User{"ak" , "akash@gmail.com" , 34}
        fmt.Printf("%v\n" ,ak)

        fmt.Printf("%+v\n", ak)

        fmt.Printf("%v\n", ak-Email)

    }
