package main

import "fmt"

func printOS() {
	osName := GetOS()
	fmt.Println("Hello I'm running on", osName)
}

func main() {
	printOS()
}
