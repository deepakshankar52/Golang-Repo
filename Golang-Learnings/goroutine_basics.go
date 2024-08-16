// package main

// import (
//     "fmt"
//     "time"
// )

// func f(from string) {
//     for i := 0; i < 3; i++ {
//         fmt.Println(from, ":", i)
//     }
// }

// func main() {

//     f("direct")

//     go f("goroutine")

//     go func(msg string) {
//         fmt.Println(msg)
//     }("going")

//     time.Sleep(time.Second)
//     fmt.Println("done")
// }

// package main
// import (
//     "fmt"
//     "time"
// )

// // Prints numbers from 1-3 along with the passed string
// func foo(s string) {
//     for i := 1; i <= 5; i++ {
//         time.Sleep(100 * time.Millisecond)
//         fmt.Println(s, ": ", i)
//     }
// }

// func main() {
    
//     // Starting two goroutines
//     go foo("1st goroutine")
//     go foo("2nd goroutine")

//     // Wait for goroutines to finish before main goroutine ends
//     time.Sleep(time.Second)
//     fmt.Println("Main goroutine finished")
// }

// package main

// import (
// 	"fmt"
// )

// func hello() {
// 	fmt.Println("Hello world goroutine")
// }
// func main() {
// 	go hello()
// 	fmt.Println("main function")
// }


// package main

// import (  
//     "fmt"
//     "time"
// )

// func hello() {  
//     fmt.Println("Hello world goroutine")
// }
// func main() {  
//     go hello()
//     time.Sleep(1 * time.Second)
//     fmt.Println("main function")
// }


package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
