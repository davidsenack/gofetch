// Functions to gather CPU information like model, cores, and usage.
package cpu

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GetCPUModel returns the model name of the CPU.
func GetCPUModel() (string, error) {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "", fmt.Errorf("error reading /proc/cpuinfo: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("model name not found in /proc/cpuinfo")
}
