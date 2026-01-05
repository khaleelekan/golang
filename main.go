package main


import "fmt"

func main() {
	var greeting string = "Hello, Go backend world!"
	var name = "alkhaleely"
	var age int = 30

	fmt.Println("greeting:", name, greeting, "you are" ,age ,"years old")
}

// func stringTypes() {
// 	var str1 string = "Hello, World!"
// }

func zeroValues() {
    var a int     // 0
    var b float64 // 0.0
    var c string  // "" (empty string)
    var d bool    // false
    // var e *int    // nil (pointer)
    // var f []int   // nil (slice)
    // var g map[string]int // nil (map)
    
    fmt.Println("int zero:", a)
    fmt.Println("float zero:", b)
    fmt.Println("string zero:", c) // Will print nothing
    fmt.Println("bool zero:", d)
}