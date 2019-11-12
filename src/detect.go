package main

import (
	"fmt"
	"os"
)

type stringMap map[string]string

// SprintRequired : prints list of software required for APBS-REST to be properly installed
func SprintRequired(required_software_list []string) string {
	var output string
	output = fmt.Sprintf("%sRequired software:\n", output)
	for _, name := range required_software_list {
		output = fmt.Sprintf("%s  - %-10s\n", output, name)
	}
	output = fmt.Sprintf("%s\n", output)

	return output
}

// SprintInstalled : prints list of existing, previously-installed software required for APBS-REST
func SprintInstalled(existing_software_list []stringMap) string {
	var output string
	output = fmt.Sprintf("%sInstalled software...\n", output)
	for _, pathMap := range existing_software_list {
		name := pathMap["name"]
		path := pathMap["path"]
		output = fmt.Sprintf("%s  - %-10s : %s\n", output, name, path)
	}
	output = fmt.Sprintf("%s\n", output)

	return output
}

// SprintNeeded : prints list of required software which isn't currently installed
func SprintNeeded(needed_software_list []stringMap) string {
	var output string
	output = fmt.Sprintf("%sNeeded software...\n", output)
	for _, swMap := range needed_software_list {
		name := swMap["name"]
		url := swMap["url"]
		output = fmt.Sprintf("%s  - %-10s - get from %s\n", output, name, url)
	}

	return output
}

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
		return 1
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
	fmt.Println(recommendation)
	return 0
}

func main() {
	os.Exit(run())
}
