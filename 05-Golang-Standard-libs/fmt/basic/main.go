package main

import (
	"fmt"
)

var (
	name   string
	age    int
	gender string
)

func main() {
	// Print Printf Println
	fmt.Println("Input args one:")
	fmt.Scan(&name, &age, &gender)
	fmt.Print("Userinfo: ")
	fmt.Printf("Name-%s, Age-%d, Gender-%s\n", name, age, gender)
	fmt.Println("---- a new line ----")
	// Scan Scanf Scanln
	fmt.Println("Input args two:")
	fmt.Scanf("1-%s, 2-%d, 3-%s\n", &name, &age, &gender)
	fmt.Print("Userinfo: ")
	fmt.Printf("Name-%s, Age-%d, Gender-%s\n", name, age, gender)
	fmt.Println("---- a new line ----")

	fmt.Println("Input args three:")
	fmt.Scanln(&name, &age, &gender) // 类似 Scan
	fmt.Print("Userinfo: ")
	fmt.Printf("Name-%s, Age-%d, Gender-%s\n", name, age, gender)
	fmt.Println("---- a new line ----")
}
