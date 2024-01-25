package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func grabBanner(host string, port int) string {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return fmt.Sprintf("Connection failed: %v", err)
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "No banner received"
	}
	return string(buf[:n])
}

func main() {
	host := flag.String("host", "127.0.0.1", "Target host")
	port := flag.Int("port", 80, "Target port")
	flag.Parse()
	fmt.Printf("[*] Banner grab: %s:%d\n", *host, *port)
	banner := grabBanner(*host, *port)
	fmt.Printf("[+] Banner:\n%s\n", banner)
}
