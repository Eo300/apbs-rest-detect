// +build windows

package main

import "os/exec"

import "strings"

import "fmt"

// CheckErr : performs obligatory error check.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetOS : returns name of operating system
func GetOS() string {
	return "Windows"
}

func Where(shellCommand string) (string, error) {
	cmdObj := exec.Command("where", shellCommand)
	stdOutBytes, err := cmdObj.Output()

	returnStr := strings.TrimSpace(string(stdOutBytes))
	return returnStr, err
}

// HasVirtHardware : checks that CPU supports Intel VT-x or AMD SVM virtualization
func HasVirtHardware() bool {
	return true
}

// HasDocker : checks if Docker is installed and is found in PATH
func HasDocker() bool {
	whereDocker, err := Where("docker")

	if len(whereDocker) == 0 && err != nil {
		return false
	}
	return true
}

// GetInstallRecommendations : check installed software and build recommended install path for user
func GetInstallRecommendations() string {
	var output string

	// list of software name:link pairings
	var method_name string
	var required_software_list []string
	var needed_software_list []stringMap
	var existing_software_list []stringMap

	if HasDocker() {
		method_name = "Docker for Desktop (w/ Kubernetes)"
		required_software_list = append(required_software_list, "Docker for Desktop")
	}

	// Print the recommended prerequisite install path
	output = fmt.Sprintf("%sRecommended Path:\n  %s\n\n", output, method_name)

	// Print list of required software to the string variable 'output'
	if len(required_software_list) > 0 {
		output = fmt.Sprintf("%s%s", output, SprintRequired(required_software_list))
	}

	// Print list of existing software
	if len(existing_software_list) > 0 {
		output = fmt.Sprintf("%s%s", output, SprintInstalled(existing_software_list))
	}

	// Print list of software needed by user
	if len(needed_software_list) > 0 {
		output = fmt.Sprintf("%s%s", output, SprintNeeded(needed_software_list))
	} else {
		output = fmt.Sprintf("%sNice, you have all the prequisite software! You're good to go install APBS-REST.", output)
	}

	return output
}

// func main() {
// 	fmt.Println("Hello I'm in Windows")
// }
