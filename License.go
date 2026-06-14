package main

import (
    "fmt"
)

func main() {
    var word string
    var num int
    
    fmt.Print("Enter license number: ")
    fmt.Scan(&num)
    
    fmt.Print("Enter plan type (PLATINUM/DIAMOND/MASTER): ")
    fmt.Scan(&word)
    
    switch word {
    case "PLATINUM":
        fmt.Println("\n# PLAN ONE PLATINUM")
        fmt.Println("Price: $100")
        fmt.Println("Key: One Month License")
        fmt.Printf("The number is: %d\n", num)
        
    case "DIAMOND":
        fmt.Println("\n# PLAN TWO DIAMOND")
        fmt.Println("Price: $200")
        fmt.Println("Key: One Year License")
        fmt.Printf("The number is: %d\n", num)
        
    case "MASTER":
        fmt.Println("\n# PLAN THREE MASTER")
        fmt.Println("Price: $500")
        fmt.Println("Key: Life Time License")
        fmt.Printf("The number is: %d\n", num)
        
    default:
        fmt.Println("Invalid plan type")
    }
}
