// Package cmd contains the main entry point for the Gofetch CLI.
package cmd

import (
	"fmt"

	"github.com/davidsenack/gofetch/pkg/cpu"
)

func main() {
	model, err := cpu.GetCPUModel()
	if err != nil {
		fmt.Println("Error fetching CPU model:", err)
		return
	}
	fmt.Println("CPU Model:", model)
}
