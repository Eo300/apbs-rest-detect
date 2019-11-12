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

	return string(stdOutput), err
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

// GetInstallRecommendations : check installed software and build recommended install path for user
func GetInstallRecommendations() string {
	var output string

	// list of software name:link pairings
	var method_name string
	var software_list []map[string]string
	var existing_software_list []map[string]string
	minikube_map := map[string]string{"name": "Minikube", "url": "https://kubernetes.io/docs/tasks/tools/install-minikube/"}
	virtualbox_map := map[string]string{"name": "Virtualbox", "url": "https://www.virtualbox.org/wiki/Downloads"}

	if HasKVM() {
		// recommend minikube via KVM
		method_name = "Minikube via KVM"
		software_list = append(software_list, minikube_map)

	} else if HasVirtualbox() {
		// recommend Minikube via Virtualbox
		method_name = "Minikube via Virtualbox"
		software_list = append(software_list, minikube_map)

		vbox_path, _ := Which("virtualbox")
		vbox_path_map := map[string]string{"name": "Virtualbox", "url": vbox_path}
		existing_software_list = append(existing_software_list, vbox_path_map)

	} else {
		method_name = "Minikube via Virtualbox"
		software_list = append(software_list, minikube_map)
		software_list = append(software_list, virtualbox_map)
	}

	output = fmt.Sprintf("%sRecommended Path:\n  %s\n\n", output, method_name)

	if len(existing_software_list) > 0 {
		output = fmt.Sprintf("%sInstalled software...\n", output)
		for _, pathMap := range existing_software_list {
			name := pathMap["name"]
			path := pathMap["path"]
			output = fmt.Sprintf("%s  %-11s : %s\n", output, name, path)
		}
		fmt.Println()
	}

	output = fmt.Sprintf("%sNeeded software...\n", output)
	for _, swMap := range software_list {
		name := swMap["name"]
		url := swMap["url"]
		output = fmt.Sprintf("%s  %-11s - get from %s\n", output, name, url)
	}

	return output
}

// func main() {
// 	// fmt.Println("Hello I'm in Linux")
// 	fmt.Println(HasVirtHardware())
// }
