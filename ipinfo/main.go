package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func lookup(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var info IPInfo
	json.NewDecoder(resp.Body).Decode(&info)
	return &info, nil
}

func main() {
	ip := flag.String("ip", "", "IP address to lookup")
	flag.Parse()
	if *ip == "" {
		fmt.Fprintln(os.Stderr, "Usage: ipinfo -ip <address>")
		os.Exit(1)
	}
	info, err := lookup(*ip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("IP:       %s\nCity:     %s\nCountry:  %s\nOrg:      %s\nTimezone: %s\n",
		info.IP, info.City, info.Country, info.Org, info.Timezone)
}
