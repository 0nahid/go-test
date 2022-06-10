// a program for showing just how to use the Println function
package main

import "fmt"

func main() {
	/**
	** int 8 → 2 ^8 = 256 = -128 0 127
	** int 16 → 2 ^16 = 65536 = -32768 0 32767
	** unit8 → 2 ^8 = 256 (0-256)
	** "Nahid Hassan" → String data type
	** 123 → int data type
	** true → bool data type
	** 0.123 → float data type
	** []int → slice data type
	** map[string]int → map data type
	** struct → struct data type
	** func → func data type
	** interface → interface data type
	** chan int → chan data type
	 */
	// variable declaration
	// var firstName, lastName string
	// var age int
	// variable assignment
	firstName := "Nahid"
	lastName := "Hassan Bulbul"
	age := 23

	// fmt.Println("Hello, World! \nThis is just a test")
	// fmt.Println("Hello I am ", firstName, lastName, " and I am ", age, " years old.")
	// fmt.Println("I am from CSE \tDepartment")
	// fmt.Println("My Identity Number id: 22103166")
	// fmt.Println("My First Name is: ", firstName)
	// fmt.Println("My Last Name is: ", lastName)
	// fmt.Println("My Age is: ", age)
	// fmt.Println(hobbies[] + " " + isMarried)
	fmt.Printf("My First Name is: %s%s and age is %d \n", firstName, lastName,age)
}
