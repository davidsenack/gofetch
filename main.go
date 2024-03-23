// Package cmd contains the main entry point for the Gofetch CLI.
package main

import (
	"fmt"
	"strings"

	"github.com/davidsenack/gofetch/pkg/cpu"
	"github.com/davidsenack/gofetch/pkg/distro"
	"github.com/davidsenack/gofetch/pkg/gpu"
	"github.com/davidsenack/gofetch/pkg/memory"
)

func main() {
	const (
		redColor   = "\033[31m"
		resetColor = "\033[0m"
	)

	printError := func(msg string, err error) {
		fmt.Printf("\n%sError fetching %s: %v%s\n", redColor, msg, err, resetColor)
	}

	printInfo := func(label, info string) {
		fmt.Printf("%s%s:%s %s\n", redColor, label, resetColor, info)
	}

	printHeader := func(header string) {
		fmt.Printf("\n%s%s%s\n", redColor, header, resetColor)
	}

	hostname, err := distro.GetUserHost()
	if err != nil {
		printError("user@hostname", err)
		return
	}
	printHeader(hostname)
	fmt.Println(strings.Repeat("-", len(hostname)+2))

	model, err := cpu.GetCPUModel()
	if err != nil {
		printError("CPU model", err)
		return
	}
	printInfo("CPU Model", model)

	kernel, err := distro.GetKernelVersion()
	if err != nil {
		printError("kernel version", err)
		return
	}
	printInfo("Kernel Version", kernel)

	currentOS, err := distro.GetDistro()
	if err != nil {
		printError("distro", err)
		return
	}
	printInfo("Distro", currentOS)

	uptime, err := distro.GetSystemUptime()
	if err != nil {
		printError("system uptime", err)
		return
	}
	printInfo("System Uptime", uptime)

	packages, packageManager, err := distro.GetInstalledPackages()
	if err != nil {
		printError("installed packages", err)
		return
	}
	printInfo("Packages", fmt.Sprintf("%d (%s)", packages, packageManager))

	shellVersion, err := distro.GetShellVersion()
	if err != nil {
		printError("shell version", err)
		return
	}
	printInfo("Shell", shellVersion)

	terminal, err := distro.GetTerminal()
	if err != nil {
		printError("terminal", err)
		return
	}
	printInfo("Terminal", terminal)

	gpu, err := gpu.GetCurrentGPU()
	if err != nil {
		printError("GPU", err)
		return
	}
	printInfo("GPU", gpu)

	currentMemory, err := memory.GetCurrentMemory()
	if err != nil {
		printError("memory", err)
		return
	}

	totalMemory, err := memory.GetTotalMemory()
	if err != nil {
		printError("memory", err)
		return
	}
	printInfo("Memory", fmt.Sprintf("%d MB / %d MB", currentMemory, totalMemory))

	distro.DisplayTerminalColors()
}
