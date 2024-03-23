// Functions to determine the operating system/distribution name, version, etc.
package distro

import (
	"fmt"
	"os"
	"strings"
)

func GetDistro() (string, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", fmt.Errorf("error reading /etc/os-release: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				return strings.Trim(parts[1], "\""), nil
			}
		}
	}

	return "", fmt.Errorf("PRETTY_NAME not found in /etc/os-release")
}

// GetKernelVersion returns the current running kernel version.
func GetKernelVersion() (string, error) {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return "", fmt.Errorf("error reading /proc/version: %w", err)
	}

	versionInfo := strings.Fields(string(data))
	if len(versionInfo) > 2 {
		return versionInfo[2], nil // Typically, the third field in /proc/version is the kernel version.
	}

	return "", fmt.Errorf("kernel version not found in /proc/version")
}
