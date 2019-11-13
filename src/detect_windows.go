// +build windows

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var systeminfo, _ = GetSystemInfo()

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

// Where :
func Where(shellCommand string) (string, error) {
	cmdObj := exec.Command("where", shellCommand)
	stdOutBytes, err := cmdObj.Output()

	returnStr := strings.TrimSpace(string(stdOutBytes))
	return returnStr, err
}

// GetSystemInfo :
func GetSystemInfo() (string, error) {
	cmdObj := exec.Command("systeminfo")
	stdOutBytes, err := cmdObj.Output()

	sysinfoStr := strings.TrimSpace(string(stdOutBytes))
	return sysinfoStr, err
}

// GetWindowsEdition :
func GetWindowsEdition(systeminfo_text string) string {
	var edition string

	scanner := bufio.NewScanner(strings.NewReader(systeminfo_text))
	for scanner.Scan() {
		splitLine := strings.SplitN(scanner.Text(), ":", 2)
		if strings.TrimSpace(splitLine[0]) == "OS Name" {
			edition = strings.TrimSpace(splitLine[1])
			break
		}
	}

	return edition
}

// HasVirtHardware : checks that CPU supports Intel VT-x or AMD SVM virtualization
func HasVirtHardware() bool {
	// TODO: Do we return True if Hyper-V is turned on OR if virtualization if available
	var systeminfoText string
	systeminfoText = systeminfo
	if strings.Contains(systeminfoText, "A hypervisor has been detected. Features required for Hyper-V will not be displayed.") {
		return true
	} else if strings.Contains(systeminfoText, "Virtualization Enabled In Firmware: Yes") {
		return true
	} else if strings.Contains(systeminfoText, "Hyper-V Requirements:") {
		return true
	}
	return false
}

// HasDocker : checks if Docker is installed and is found in PATH
func HasDocker() bool {
	whereDocker, err := Where("docker")

	if len(whereDocker) == 0 && err != nil {
		return false
	}
	return true
}

// HasVirtualbox : checks if Virtualbox is installed and found in systeminfo
func HasVirtualbox(systeminfo_text string) bool {
	if !strings.Contains(systeminfo_text, "VirtualBox") {
		return false
	}
	return true
}

// HasMinikube : checks if Minikube is installed and is found in PATH
func HasMinikube() bool {
	whichOut, err := Where("minikube")

	if len(whichOut) == 0 && err != nil {
		return false
	}
	return true
}

// GetInstallRecommendations : check installed software and build recommended install path for user
func GetInstallRecommendations() string {
	var output string
	// var systeminfo string
	var windowsEdition string
	var validEditions []string

	// list of software name:link pairings
	var method_name string
	var required_software_list []string
	var missing_software_list []stringMap
	var existing_software_list []stringMap
	docker_map := stringMap{"name": "Docker Desktop", "url": "https://docs.docker.com/docker-for-windows/install/"}
	minikube_map := stringMap{"name": "Minikube", "url": "https://kubernetes.io/docs/tasks/tools/install-minikube/"}
	virtualbox_map := stringMap{"name": "VirtualBox", "url": "https://www.virtualbox.org/wiki/Downloads"}

	// systeminfo, _ = GetSystemInfo()
	windowsEdition = GetWindowsEdition(systeminfo)
	validEditions = []string{"Pro", "Professional", "Enterprise", "Education"}
	// validEditions = []string{"Pro", "Professional", "Education"}

	if HasDocker() {
		// recommend Docker Desktop
		method_name = "Docker Desktop (w/ Kubernetes)"
		required_software_list = append(required_software_list, "Docker Desktop")

		docker_path, _ := Where("docker")
		docker_path_map := stringMap{"name": "Docker", "path": docker_path}
		existing_software_list = append(existing_software_list, docker_path_map)
	} else if HasVirtualbox(systeminfo) {
		// recommend Minikube via Virtualbox
		method_name = "Minikube (via VirtualBox)"
		required_software_list = append(required_software_list, "VirtualBox", "Minikube")

		// get VirtualBox path
		expectedVBoxPath := "C:\\Program Files\\Oracle\\VirtualBox\\VirtualBox.exe"
		_, err := os.Stat(expectedVBoxPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not find VirtualBox in expected path: %s\n", expectedVBoxPath)
			// missing_software_list = append(missing_software_list, virtualbox_map)
		} else {
			vBox_path_map := stringMap{"name": "VirtualBox", "path": expectedVBoxPath}
			existing_software_list = append(existing_software_list, vBox_path_map)
		}

		// check for minikube in PATH
		if HasMinikube() {
			minikube_path, _ := Where("minikube")
			minikube_path_map := stringMap{"name": "Minikube", "path": minikube_path}
			existing_software_list = append(existing_software_list, minikube_path_map)
		} else {
			missing_software_list = append(missing_software_list, minikube_map)
		}

	} else {
		fields := strings.Fields(windowsEdition)
		lastWord := fields[len(fields)-1] // should be "Home", "Pro", "Enterprise", or "Education"

		if InSliceString(validEditions, lastWord) { // If Windows Professional or above
			// recommend Docker Desktop
			method_name = "Docker Desktop (w/ Kubernetes)"
			required_software_list = append(required_software_list, "Docker Desktop")
			missing_software_list = append(missing_software_list, docker_map)
		} else {
			// recommend Minikube via Virtualbox
			method_name = "Minikube (via VirtualBox)"
			required_software_list = append(required_software_list, "VirtualBox", "Minikube")
			missing_software_list = append(missing_software_list, virtualbox_map)
			missing_software_list = append(missing_software_list, minikube_map)
		}

	}

	// Print the recommended prerequisite install path
	output = fmt.Sprintf("%sRecommended Path:\n  %s\n\n", output, method_name)

	// Print the user's Windows edition name
	output = fmt.Sprintf("%sWindows Edition:\n  %s\n\n", output, windowsEdition)

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
