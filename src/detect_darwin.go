// +build darwin

package main

// CheckErr : performs obligatory error check.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetOS : returns name of operating system
func GetOS() string {
	return "macOS"
}

// HasVirtHardware : checks that CPU supports Intel VT-x or AMD SVM virtualization
func HasVirtHardware() bool {
	return true
}

// GetInstallRecommendations : check installed software and build recommended install path for user
func GetInstallRecommendations() string {
	var output string
	output = ""

	return output
}

// func main() {
// 	fmt.Println("Hello I'm in Linux")
// }
