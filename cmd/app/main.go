package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"steam-lancache/internal/api"
	"steam-lancache/internal/config"
	"steam-lancache/internal/dns"
	"steam-lancache/internal/helpers"
	"steam-lancache/internal/proxy"
	"steam-lancache/internal/stats"
	"steam-lancache/internal/tcp"
)

func main() {
	// init log
	logfile, err := os.Create("application.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	multiwriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiwriter)
	log.SetFlags(0)

	// display system info
	fmt.Println("[+] Checking...")
	_, _, _, _, hasInternet, _ := helpers.ShowSystemInfo(true)
	if !hasInternet {
		log.Fatal("[!] No Internet Connection")
	}

	// get token from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[?] Enter your Token: ")
	tokenInput, _ := reader.ReadString('\n')
	token := strings.TrimSpace(tokenInput)

	if token == "" {
		log.Fatal("[!] Token is required")
	}

	// get token info
	log.Println("[+] Receiving token information...")
	ip, details, err := api.GetTokenInfo(token)
	if err != nil {
		log.Fatalf("[!] Receiving token information failed -> %v", err)
	}

	// show details
	fmt.Println("------------------------------------------------")
	fmt.Printf("[*] Subscription Active\n")
	fmt.Printf("[*] User IP:     %s\n", details.UserIP)
	fmt.Printf("[*] Expires:     %s\n", details.End)
	fmt.Printf("[*] Upstream IP: %s\n", ip)
	fmt.Println("------------------------------------------------")

	// fetch and select DNS server
	log.Println("\nFetching DNS servers...")
	dnsServers, err := dns.FetchDNSServers()
	if err != nil {
		log.Fatalf("[!] Failed to fetch DNS servers -> %v", err)
	}

	selectedDNS := dns.DNSSelection(dnsServers)

	fmt.Printf("\n[+] Selected DNS: %s -> %s\n", selectedDNS.Name, selectedDNS.IP)

	appState := &config.AppState{
		UserToken:     token,
		CacheServerIP: ip,
		DNSIP:         selectedDNS.IP,
	}

	// init traffic stat service
	stats.Load()

	// start services
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		stats.StartSaver()
		wg.Done()
	}()

	go func() {
		dns.Start(appState)
		wg.Done()
	}()

	go func() {
		tcp.Start(appState.CacheServerIP)
		wg.Done()
	}()

	proxy.Start(appState)
}
