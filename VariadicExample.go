
package main

import "fmt"

// Sum returns the sum of adding all the numbers in
// the parameters together.
func Sum(numbers ...int) int {

    n := 0

    for _,number := range numbers {
        n += number
    }

    return n
}

func main() {
    sm1 := Sum(1, 2, 3, 4) // = 1 + 2 + 3 + 4 = 10

    sm2 := Sum(1, 2) // = 1 + 2 = 3

    sm3 := Sum(7, 1, -2, 0, 18) // = 7 + 1 + -2 + 0 + 18 = 24

    sm4 := Sum() // = 0


    fmt.Printf("sm1 = %d\n", sm1)
    fmt.Printf("sm2 = %d\n", sm2)
    fmt.Printf("sm3 = %d\n", sm3)
    fmt.Printf("sm4 = %d\n", sm4)

    // Output:
    // sm1 = 10
    // sm2 = 3
    // sm3 = 24
    // sm4 = 0
}
