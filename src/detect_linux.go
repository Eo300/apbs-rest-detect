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

// HasKVM : checks if KVM is install and is found in PATH
func HasKVM() bool {
	whichOut, err := Which("kvm")

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
	var needed_software_list []map[string]string
	var existing_software_list []map[string]string
	minikube_map := map[string]string{"name": "Minikube", "url": "https://kubernetes.io/docs/tasks/tools/install-minikube/"}
	virtualbox_map := map[string]string{"name": "VirtualBox", "url": "https://www.virtualbox.org/wiki/Downloads"}

	if HasKVM() {
		// recommend minikube via KVM
		method_name = "Minikube (via KVM)"
		required_software_list = append(required_software_list, "KVM", "Minikube")

	} else if HasVirtualbox() {
		// recommend Minikube via Virtualbox
		method_name = "Minikube (via VirtualBox)"
		required_software_list = append(required_software_list, "VirtualBox", "Minikube")
		// needed_software_list = append(needed_software_list, minikube_map)

		vbox_path, _ := Which("virtualbox")
		vbox_path_map := map[string]string{"name": "Virtualbox", "path": vbox_path}
		existing_software_list = append(existing_software_list, vbox_path_map)

	} else {
		// recommend Minikube via Virtualbox
		method_name = "Minikube (via VirtualBox)"
		required_software_list = append(required_software_list, "VirtualBox", "Minikube")
		// needed_software_list = append(needed_software_list, minikube_map)
		needed_software_list = append(needed_software_list, virtualbox_map)
	}

	// check for minikube in PATH
	if HasMinikube(){
		minikube_path, _ := Which("minikube")
		minikube_path_map := map[string]string{"name": "Minikube", "path": minikube_path}
		existing_software_list = append(existing_software_list, minikube_path_map)
	} else{
		needed_software_list = append(needed_software_list, minikube_map)
	}

	output = fmt.Sprintf("%sRecommended Path:\n  %s\n\n", output, method_name)

	// Print list of required software
	if len(required_software_list) > 0 {
		output = fmt.Sprintf("%sRequired software:\n", output)
		for _, name := range required_software_list {
			output = fmt.Sprintf("%s  - %-10s\n", output, name)
		}
		output = fmt.Sprintf("%s\n", output)
	}
	
	// Print list of existing software
	if len(existing_software_list) > 0 {
		output = fmt.Sprintf("%sInstalled software...\n", output)
		for _, pathMap := range existing_software_list {
			name := pathMap["name"]
			path := pathMap["path"]
			output = fmt.Sprintf("%s  - %-10s : %s\n", output, name, path)
		}
		output = fmt.Sprintf("%s\n", output)
	}

	// Print list of software needed by user
	if len(needed_software_list) > 0 {
		output = fmt.Sprintf("%sNeeded software...\n", output)
		for _, swMap := range needed_software_list {
			name := swMap["name"]
			url := swMap["url"]
			output = fmt.Sprintf("%s  - %-10s - get from %s\n", output, name, url)
		}
	} else{
		output = fmt.Sprintf("%sNice, you have all the prequisite software! You're good to go install APBS-REST.", output)
	}

	return output
}

// func main() {
// 	// fmt.Println("Hello I'm in Linux")
// 	fmt.Println(HasVirtHardware())
// }
