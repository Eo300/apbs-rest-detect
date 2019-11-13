// +build linux

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// CheckErr : performs obligatory error check.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetOS : returns name of operating system
func GetOS() string {
	return "Linux"
}

// Which : executes "which" shell command, returns output as string
func Which(shellCommand string) (string, error) {
	// println("executing: which", shellCommand)
	cmdObj := exec.Command("which", shellCommand)
	stdOutput, err := cmdObj.Output()

	returnStr := strings.TrimSpace(string(stdOutput))
	return returnStr, err
}

// HasVirtHardware : checks that CPU supports Intel VT-x or AMD SVM virtualization
func HasVirtHardware() bool {
	// $ lscpu
	cpuCmd := exec.Command("lscpu")
	cpuOut, err := cpuCmd.Output()
	CheckErr(err)

	// $ grep "Virtualization:"
	grepCmd := exec.Command("grep", "Virtualization:")
	stdinPipe, _ := grepCmd.StdinPipe()
	stdinPipe.Write(cpuOut)
	stdinPipe.Close()

	grepOut, err := grepCmd.Output()
	output := string(grepOut)
	output = strings.TrimSpace(output)

	// Checks if there was output for grep-ing "Virtualzation"
	if len(output) == 0 && err != nil {
		return false
	}
	return true
}

// HasKVM : checks if KVM is installed and is found in PATH
func HasKVM() bool {
	whichOut, err := Which("virsh")

	if len(whichOut) == 0 && err != nil {
		return false
	}
	return true
}

// HasVirtualbox : checks if Virtualbox is installed and is found in PATH
func HasVirtualbox() bool {
	whichOut, err := Which("virtualbox")

	if len(whichOut) == 0 && err != nil {
		return false
	}
	return true
}

// HasMinikube : checks if Minikube is installed and is found in PATH
func HasMinikube() bool {
	whichOut, err := Which("minikube")

	if len(whichOut) == 0 && err != nil {
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
	var missing_software_list []stringMap
	var existing_software_list []stringMap
	minikube_map := stringMap{"name": "Minikube", "url": "https://kubernetes.io/docs/tasks/tools/install-minikube/"}
	virtualbox_map := stringMap{"name": "VirtualBox", "url": "https://www.virtualbox.org/wiki/Downloads"}

	if HasKVM() {
		// recommend minikube via KVM
		method_name = "Minikube (via KVM)"
		required_software_list = append(required_software_list, "KVM", "Minikube")

	} else if HasVirtualbox() {
		// recommend Minikube via VirtualBox
		method_name = "Minikube (via VirtualBox)"
		required_software_list = append(required_software_list, "VirtualBox", "Minikube")

		vbox_path, _ := Which("virtualbox")
		vbox_path_map := stringMap{"name": "Virtualbox", "path": vbox_path}
		existing_software_list = append(existing_software_list, vbox_path_map)

	} else {
		// recommend Minikube via Virtualbox
		method_name = "Minikube (via VirtualBox)"
		required_software_list = append(required_software_list, "VirtualBox", "Minikube")
		missing_software_list = append(missing_software_list, virtualbox_map)
	}

	// check for minikube in PATH
	if HasMinikube() {
		minikube_path, _ := Which("minikube")
		minikube_path_map := stringMap{"name": "Minikube", "path": minikube_path}
		existing_software_list = append(existing_software_list, minikube_path_map)
	} else {
		missing_software_list = append(missing_software_list, minikube_map)
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
	if len(missing_software_list) > 0 {
		output = fmt.Sprintf("%s%s", output, SprintNeeded(missing_software_list))
	} else {
		output = fmt.Sprintf("%sNice, you have all the prequisite software! You're good to go install APBS-REST.", output)
	}

	return output
}
