// Package cmd contains the main entry point for the Gofetch CLI.
package main

import (
	"fmt"

	"github.com/davidsenack/gofetch/pkg/cpu"
	"github.com/davidsenack/gofetch/pkg/distro"
)

func main() {
	model, err := cpu.GetCPUModel()
	if err != nil {
		fmt.Println("Error fetching CPU model:", err)
		return
	}
	fmt.Println("CPU Model:", model)

	kernel, err := distro.GetKernelVersion()
	if err != nil {
		fmt.Println("Error fetching kernel version:", err)
		return
	}
	fmt.Println("Kernel Version:", kernel)

	distro, err := distro.GetDistro()
	if err != nil {
		fmt.Println("Error fetching distro:", err)
		return
	}
	fmt.Println("Distro:", distro)
}
