package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Run, print PID and hang there so I can attach to it from regs.go
func main() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	select {}
}
