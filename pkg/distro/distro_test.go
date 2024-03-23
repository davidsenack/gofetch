package distro

import (
	"fmt"
	"testing"
)

func TestGetDistro(t *testing.T) {
	distro, err := GetDistro()
	if err != nil {
		t.Errorf("GetDistro() failed, expected no error, got %v", err)
	}
	if distro == "" {
		t.Errorf("GetDistro() failed, expected a non-empty string, got an empty string")
	}
}

func TestGetKernelVersion(t *testing.T) {
	version, err := GetKernelVersion()
	if err != nil {
		t.Errorf("GetKernelVersion() failed, expected no error, got %v", err)
	}
	if version == "" {
		t.Errorf("GetKernelVersion() failed, expected a non-empty string, got an empty string")
	}
}

func TestGetSystemUptime(t *testing.T) {
	uptime, err := GetSystemUptime()
	if err != nil {
		t.Errorf("GetSystemUptime() failed, expected no error, got %v", err)
	}
	uptimeInt, err := parseUptime(uptime) // Convert uptime from formatted string to total seconds
	if err != nil {
		t.Errorf("Failed to parse uptime, got error: %v", err)
	}
	if uptimeInt <= 0 {
		t.Errorf("GetSystemUptime() failed, expected a positive uptime, got %d", uptimeInt)
	}
}

func TestGetInstalledPackages(t *testing.T) {
	packages, packageManager, err := GetInstalledPackages()
	if err != nil {
		t.Errorf("GetInstalledPackages() failed, expected no error, got %v", err)
	}
	if packages < 0 {
		t.Errorf("GetInstalledPackages() failed, expected a non-negative number of packages, got %d", packages)
	}
	if packageManager == "" {
		t.Errorf("GetInstalledPackages() failed, expected a non-empty package manager, got an empty string")
	}
}

func TestGetShellVersion(t *testing.T) {
	shellVersion, err := GetShellVersion()
	if err != nil {
		t.Errorf("GetShellVersion() failed, expected no error, got %v", err)
	}
	if shellVersion == "" {
		t.Errorf("GetShellVersion() failed, expected a non-empty string, got an empty string")
	}
}

func TestGetTerminal(t *testing.T) {
	terminal, err := GetTerminal()
	if err != nil {
		t.Errorf("GetTerminal() failed, expected no error, got %v", err)
	}
	if terminal == "" {
		t.Errorf("GetTerminal() failed, expected a non-empty string, got an empty string")
	}
}

// parseUptime parses a formatted uptime string (e.g., "1d 2h 3m 4s") into total seconds.
func parseUptime(uptime string) (int, error) {
	var days, hours, minutes, seconds int
	_, err := fmt.Sscanf(uptime, "%dd %dh %dm %ds", &days, &hours, &minutes, &seconds)
	if err != nil {
		return 0, err
	}
	totalSeconds := days*86400 + hours*3600 + minutes*60 + seconds
	return totalSeconds, nil
}
