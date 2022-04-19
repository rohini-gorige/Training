package main

import "fmt"

func main() {
	var fname string
	fmt.Println("Enter your Firstname: ")
	fmt.Scan(&fname)
	fmt.Println("Enter your Lastname: ")
	var lname string
	fmt.Scan(&lname)
	fmt.Println("Your Full Name is " + fname + " " + lname)
}
