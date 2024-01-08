package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(host string, port int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err == nil {
		conn.Close()
		results <- port
	}
}

func main() {
	host := flag.String("host", "127.0.0.1", "Target host")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")
	flag.Parse()

	results := make(chan int, 100)
	var wg sync.WaitGroup

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go scanPort(*host, port, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("Scanning %s ports %d-%d\n", *host, *startPort, *endPort)
	for port := range results {
		fmt.Printf("  [OPEN] port %d\n", port)
	}
}
