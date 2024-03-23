// Functions to gather memory usage, total memory, and swap information.
package memory

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetTotalMemory returns the total memory available in MB.
func GetTotalMemory() (int, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("error reading /proc/meminfo: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				totalKb, err := strconv.Atoi(parts[1])
				if err != nil {
					return 0, fmt.Errorf("error converting MemTotal to int: %w", err)
				}
				return totalKb / 1024, nil // Convert KB to MB
			}
		}
	}

	return 0, fmt.Errorf("MemTotal not found in /proc/meminfo")
}

// GetCurrentMemory returns the current used memory in MB.
func GetCurrentMemory() (int, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("error reading /proc/meminfo: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var totalKb, freeKb, buffersKb, cachedKb int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 {
			switch parts[0] {
			case "MemTotal:":
				totalKb, err = strconv.Atoi(parts[1])
				if err != nil {
					return 0, fmt.Errorf("error converting MemTotal to int: %w", err)
				}
			case "MemFree:":
				freeKb, err = strconv.Atoi(parts[1])
				if err != nil {
					return 0, fmt.Errorf("error converting MemFree to int: %w", err)
				}
			case "Buffers:":
				buffersKb, err = strconv.Atoi(parts[1])
				if err != nil {
					return 0, fmt.Errorf("error converting Buffers to int: %w", err)
				}
			case "Cached:":
				cachedKb, err = strconv.Atoi(parts[1])
				if err != nil {
					return 0, fmt.Errorf("error converting Cached to int: %w", err)
				}
			}
		}
	}

	if totalKb == 0 {
		return 0, fmt.Errorf("MemTotal not found in /proc/meminfo")
	}

	// Adjusting the calculation to consider buffers and cached memory as free memory
	adjustedFreeKb := freeKb + buffersKb + cachedKb
	usedKb := totalKb - adjustedFreeKb
	return usedKb / 1024, nil // Convert KB to MB
}
