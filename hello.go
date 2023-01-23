// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Hello, Surya Venkatesh Vijjana.")
// 	fmt.Println(os.Args[1:len(os.Args)])
// }
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    in := []int{2, 5, 6}
    randomIndex := rand.Intn(len(in))
    pick := in[randomIndex]
    fmt.Println(pick)
}