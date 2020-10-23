package main 

import "fmt"

func main () {
    arr := [7]string{"This" , "is" , "the" , "tutorial" , "of" , "Go" , "Language"}
    fmt.Println("Array:", arr)
    myslice := arr[1:6]
    fmt.Println("Slice: ", myslice)
    fmt.Printf("Length of the slice : %d", cap(myslice))

}