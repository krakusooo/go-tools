package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "Log file path")
	filter := flag.String("filter", "", "Filter string")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if *filter == "" || strings.Contains(line, *filter) {
			fmt.Println(line)
			count++
		}
	}
	fmt.Fprintf(os.Stderr, "\n--- %d lines matched ---\n", count)
}
