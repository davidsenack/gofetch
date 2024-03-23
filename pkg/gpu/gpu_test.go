package gpu

import (
	"testing"
)

func TestGetCurrentGPU(t *testing.T) {
	gpu, err := GetCurrentGPU()
	if err != nil {
		t.Errorf("GetCurrentGPU() failed, expected no error, got %v", err)
	}
	if gpu == "" {
		t.Errorf("GetCurrentGPU() failed, expected a non-empty string, got an empty string")
	}
}
