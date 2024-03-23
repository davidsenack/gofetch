package distro

import (
	"strconv" // Added for string to int conversion
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
	uptimeInt, err := strconv.Atoi(uptime) // Convert uptime from string to int
	if err != nil {
		t.Errorf("Failed to convert uptime to integer, got error: %v", err)
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
