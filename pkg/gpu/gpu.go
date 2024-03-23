package gpu

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetCurrentGPU returns the model name of the current running GPU.
func GetCurrentGPU() (string, error) {
	cmd := exec.Command("lshw", "-class", "display")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing lshw command: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		if strings.Contains(line, "product:") {
			// Assuming the product line contains the GPU model name
			return strings.TrimSpace(strings.SplitN(line, "product:", 2)[1]), nil
		}
		if strings.Contains(line, "description:") && i+1 < len(lines) && strings.Contains(lines[i+1], "product:") {
			// If the product line immediately follows a description line, use it
			return strings.TrimSpace(strings.SplitN(lines[i+1], "product:", 2)[1]), nil
		}
	}

	return "", fmt.Errorf("GPU model name not found")
}
