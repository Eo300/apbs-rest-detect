package main

import (
	"fmt"
	"os"
)

func printOS() {
	osName := GetOS()
	fmt.Println("Hello I'm running on", osName)
}

func run() int {
	// Print the current os based on the build
	printOS()
	osName := GetOS()

	// Does the user's CPU support virtualization
	if !HasVirtHardware() {
		fmt.Fprintln(os.Stderr, "Unfortunately, your CPU does not support virtualization.")
		return 1
	}

	if osName == "Linux" {

	} else if osName == "Windows" {

	} else if osName == "macOS" {

	} else {
		fmt.Fprintln(os.Stderr, "Target OS of this build is not supported.")
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
