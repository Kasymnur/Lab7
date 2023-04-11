package main

import "fmt"

func main() {
    // sample input list
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // create a new list to store the divisible numbers
    divisibleBy3 := []int{}

    // iterate over the input list and append divisible numbers to the new list
    for _, num := range nums {
        if num%3 == 0 {
            divisibleBy3 = append(divisibleBy3, num)
        }
    }

    // print the new list
    fmt.Println(divisibleBy3)
}
