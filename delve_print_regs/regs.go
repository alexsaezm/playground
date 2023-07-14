package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

const (
	PTRACE_GETREGS   = 12 // ptrace request to get register values
	PTRACE_GETFPREGS = 14 // ptrace request to get floating-point register values
)

// RegisterSet represents the register set for PPC64LE architecture
type RegisterSet struct {
	// Define the general-purpose registers
	GPR [32]uint64
	// Define the floating-point registers
	FPR [32]uint64
	// Define additional registers as needed
	PC  uint64
	MSR uint64
	LR  uint64
	CTR uint64
	XER uint64
}

// GetRegisters retrieves all the registers using ptrace for a PPC64LE machine
func GetRegisters(pid int) (*RegisterSet, error) {
	// Attach to the process using ptrace
	if err := syscall.PtraceAttach(pid); err != nil {
		return nil, fmt.Errorf("failed to attach to process: %v", err)
	}

	// Wait for the process to stop
	if _, err := syscall.Wait4(pid, nil, syscall.WALL, nil); err != nil {
		return nil, fmt.Errorf("failed to wait for process: %v", err)
	}

	// Retrieve the general-purpose register values using PTRACE_GETREGS request
	var regs RegisterSet
	if _, _, err := syscall.Syscall6(syscall.SYS_PTRACE, uintptr(PTRACE_GETREGS), uintptr(pid), uintptr(0), uintptr(unsafe.Pointer(&regs)), 0, 0); err != 0 {
		return nil, fmt.Errorf("failed to get general-purpose register values: %v", err)
	}

	return &regs, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run regs.go <PID>")
		return
	}

	pidStr := os.Args[1]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Printf("Invalid PID: %s\n", pidStr)
		return
	}

	regs, err := GetRegisters(pid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("General-Purpose Registers:\n%+v\n", regs.GPR)
	fmt.Printf("Floating-Point Registers:\n%+v\n", regs.FPR)
}
