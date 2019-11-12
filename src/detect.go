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
	osName := GetOS()
	fmt.Printf("Target: %s\n\n", osName)

	// Does the user's CPU support virtualization
	if !HasVirtHardware() {
		fmt.Fprintln(os.Stderr, "Unfortunately, your CPU does not support virtualization.")
		// return 1
	}

	var recommendation string
	if osName == "Linux" {
		recommendation = GetInstallRecommendations()
	} else if osName == "Windows" {
		recommendation = GetInstallRecommendations()
	} else if osName == "macOS" {
		recommendation = GetInstallRecommendations()
	} else {
		fmt.Fprintln(os.Stderr, "Target OS of this build is not supported.")
		return 1
	}

	// Print final recommendation to stdout
	// fmt.Println("Based on installations present on your system, we recommend the following software for satisfying the prerequisites of installing APBS-REST:")
	fmt.Println(recommendation)
	return 0
}

func main() {
	os.Exit(run())
}
