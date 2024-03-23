package cpu

import (
	"testing"
)

func TestGetCPUModel(t *testing.T) {
	model, err := GetCPUModel()
	if err != nil {
		t.Errorf("GetCPUModel() failed, expected no error, got %v", err)
	}
	if model == "" {
		t.Errorf("GetCPUModel() failed, expected a non-empty string, got an empty string")
	}
}
