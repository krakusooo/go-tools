package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func getCPUCount() int {
	return runtime.NumCPU()
}

func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func getHostname() string {
	h, _ := os.Hostname()
	return h
}

func main() {
	fmt.Println("=== System Monitor ===")
	fmt.Printf("Hostname:  %s\n", getHostname())
	fmt.Printf("OS:        %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("CPUs:      %d\n", getCPUCount())
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())

	m := getMemStats()
	fmt.Printf("Alloc:     %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc: %v MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys:       %v MB\n", m.Sys/1024/1024)
	fmt.Printf("Time:      %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
