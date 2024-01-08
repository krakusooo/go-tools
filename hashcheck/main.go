package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
)

func hashFile(path, algo string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	switch algo {
	case "md5":
		h := md5.New()
		io.Copy(h, f)
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	case "sha256":
		h := sha256.New()
		io.Copy(h, f)
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	default:
		return "", fmt.Errorf("unknown algorithm: %s", algo)
	}
}

func main() {
	filePath := flag.String("file", "", "File to hash")
	algo := flag.String("algo", "sha256", "Hash algorithm: md5 or sha256")
	flag.Parse()

	if *filePath == "" {
		fmt.Fprintln(os.Stderr, "Error: -file required")
		os.Exit(1)
	}

	hash, err := hashFile(*filePath, *algo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s  %s\n", hash, *filePath)
}
