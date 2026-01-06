package main


import ("fmt"
    "bufio"
    "os"
    "strings"
 "strconv")

func main() {
	var greeting string = "Hello, Go backend world!"
	var name = "alkhaleely"
	var age int = 30

	fmt.Println("greeting:", name, greeting, "you are" ,age ,"years old")

    var array [3]int = [3]int{1, 2, 3}
    fmt.Println("array:", array)
}

func stringTypes() {
	var str1 string = "Hello, World!"
}

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

func PersonInfo() {
    var name string = "John Doe"
    var age int = 25
    var height float64 = 5.9

    ageStr := strconv.Itoa(age)
     // Convert int to string
    const (
        daysInWeek = 7
    )
// }

func interactiveMenu() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Say Hello")
		fmt.Println("2. Calculate")
		fmt.Println("3. Exit")
		fmt.Print("Choice: ")
		
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			fmt.Print("Enter your name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Printf("Hello, %s!\n", name)
			
		case "2":
			fmt.Print("Enter first number: ")
			num1Str, _ := reader.ReadString('\n')
			num1, _ := strconv.Atoi(strings.TrimSpace(num1Str))
			
			fmt.Print("Enter second number: ")
			num2Str, _ := reader.ReadString('\n')
			num2, _ := strconv.Atoi(strings.TrimSpace(num2Str))
			
			fmt.Printf("Sum: %d\n", num1+num2)
			
		case "3":
			fmt.Println("Goodbye!")
			return
			
		default:
			fmt.Println("Invalid choice!")
		}
	}
}