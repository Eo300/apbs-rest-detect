// +build linux

package main

import (
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
	println("executing: which", shellCommand)
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
	output = ""

	// list of software name:link pairings
	// var software_list []map[string]string

	if HasKVM() {
		// recommend minikube via KVM
		// minikube_map := map[string]string{"Minikube": "https://kubernetes.io/docs/tasks/tools/install-minikube/"}
		// software_list = append(software_list, minikube_map)
	} else if HasVirtualbox() {
		// recommend Minikube via Virtualbox
	} else {

	}

	return output
}

// func main() {
// 	// fmt.Println("Hello I'm in Linux")
// 	fmt.Println(HasVirtHardware())
// }
