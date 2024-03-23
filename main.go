// Package cmd contains the main entry point for the Gofetch CLI.
package main

import (
	"fmt"

	"github.com/davidsenack/gofetch/pkg/cpu"
	"github.com/davidsenack/gofetch/pkg/distro"
	"github.com/davidsenack/gofetch/pkg/gpu"
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

	currentOS, err := distro.GetDistro()
	if err != nil {
		fmt.Println("Error fetching distro:", err)
		return
	}
	fmt.Println("Distro:", currentOS)

	uptime, err := distro.GetSystemUptime()
	if err != nil {
		fmt.Println("Error fetching system uptime:", err)
		return
	}
	fmt.Println("System Uptime:", uptime)

	packages, count, err := distro.GetInstalledPackages()
	if err != nil {
		fmt.Println("Error fetching installed packages:", err)
		return
	}
	fmt.Println("Packages:", count, "(", packages, ")")

	shellVersion, err := distro.GetShellVersion()
	if err != nil {
		fmt.Println("Error fetching shell version:", err)
		return
	}

	fmt.Printf("Shell: %s\n", shellVersion)

	terminal, err := distro.GetTerminal()
	if err != nil {
		fmt.Println("Error fetching terminal:", err)
		return
	}
	fmt.Println("Terminal:", terminal)

	gpu, err := gpu.GetCurrentGPU()
	if err != nil {
		fmt.Println("Error fetching GPU:", err)
		return
	}
	fmt.Println("GPU:", gpu)
}
