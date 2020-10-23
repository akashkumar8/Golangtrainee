package main

import (
    "fmt"
    "sort"

)

func main() {
    myints := []int{4,7,3,2,8}
    isSorted := sort.IntsAreSorted(myints)
    fmt.Println("Are ints dorted:", isSorted)
}