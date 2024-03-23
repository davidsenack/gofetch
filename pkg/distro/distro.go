// Functions to determine the operating system/distribution name, version, etc.
package distro

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
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

// GetSystemUptime returns the current system uptime in a human-readable format.
func GetSystemUptime() (string, error) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "", fmt.Errorf("error reading /proc/uptime: %w", err)
	}

	uptimeInfo := strings.Fields(string(data))
	if len(uptimeInfo) > 0 {
		uptimeSeconds, err := strconv.ParseFloat(uptimeInfo[0], 64)
		if err != nil {
			return "", fmt.Errorf("error parsing uptime: %w", err)
		}

		days := int(uptimeSeconds) / (24 * 3600)
		hours := int(uptimeSeconds) % (24 * 3600) / 3600
		minutes := int(uptimeSeconds) % 3600 / 60
		seconds := int(uptimeSeconds) % 60

		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds), nil
	}

	return "", fmt.Errorf("uptime information not found in /proc/uptime")
}

// GetInstalledPackages returns the number and type of installed packages.
func GetInstalledPackages() (int, string, error) {
	var count int
	var pkgType string

	// Check for the presence of dpkg (Debian-based systems)
	if _, err := os.Stat("/usr/bin/dpkg"); err == nil {
		data, err := exec.Command("dpkg", "--list").Output()
		if err != nil {
			return 0, "", fmt.Errorf("error listing dpkg packages: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		count = len(lines) - 6 // Adjusting for header/footer lines in dpkg output
		pkgType = "dpkg"
	} else if _, err := os.Stat("/usr/bin/rpm"); err == nil { // Check for the presence of rpm (Red Hat-based systems)
		data, err := exec.Command("rpm", "-qa").Output()
		if err != nil {
			return 0, "", fmt.Errorf("error listing rpm packages: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		count = len(lines) - 1 // Adjusting for the last empty line in rpm output
		pkgType = "rpm"
	} else {
		return 0, "", fmt.Errorf("package manager not supported")
	}

	return count, pkgType, nil
}

// GetShell returns the name of the current shell being used.
func GetShell() (string, error) {
	shellPath, ok := os.LookupEnv("SHELL")
	if !ok {
		return "", fmt.Errorf("SHELL environment variable not set")
	}
	shellParts := strings.Split(shellPath, "/")
	shell := shellParts[len(shellParts)-1] // Extract the shell name from the path

	return shell, nil
}

// GetShellVersion returns the version of the current shell being used.
func GetShellVersion() (string, error) {
	shell, err := GetShell()
	if err != nil {
		return "", err
	}

	var cmd *exec.Cmd
	switch shell {
	case "bash":
		cmd = exec.Command("bash", "--version")
	case "zsh":
		cmd = exec.Command("zsh", "--version")
	default:
		return "", fmt.Errorf("shell version retrieval not supported for %s", shell)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting shell version: %w", err)
	}

	versionInfo := strings.Split(string(output), "\n")[0] // Assuming the first line contains the version
	return versionInfo, nil
}

// GetTerminal returns the name of the current terminal being used.
func GetTerminal() (string, error) {
	terminal, ok := os.LookupEnv("TERM")
	if !ok {
		return "", fmt.Errorf("TERM environment variable not set")
	}
	return terminal, nil
}

// GetUserHost returns the current user and hostname in the format user@hostname.
func GetUserHost() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %w", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("error getting hostname: %w", err)
	}

	return fmt.Sprintf("%s@%s", user.Username, hostname), nil
}
