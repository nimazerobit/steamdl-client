package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"steam-lancache/cmd/app/tui"
	"steam-lancache/internal/api"
	"steam-lancache/internal/config"
	"steam-lancache/internal/dns"
	"steam-lancache/internal/helpers"
	"steam-lancache/internal/proxy"
	"steam-lancache/internal/stats"
	"steam-lancache/internal/tcp"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// init log
	logfile, err := os.Create("application.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	multiwriter := io.MultiWriter(logfile)
	log.SetOutput(multiwriter)
	log.SetFlags(0)

	// display system info
	fmt.Println("[+] Checking...")
	_, _, _, _, hasInternet, _ := helpers.ShowSystemInfo(true)
	if !hasInternet {
		log.Fatal("[!] No Internet Connection")
	}
	config.LoadDomains()

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
	details, err := api.GetTokenInfo(token)
	if err != nil {
		log.Fatalf("[!] Receiving token information failed -> %v", err)
	}

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
		CacheServerIP: details.UpstreamIP,
		DNSIP:         selectedDNS.IP,
	}

	// init traffic stat service
	stats.Load()

	// start services
	go func() { stats.StartSaver() }()
	go func() { dns.Start(appState) }()
	go func() { tcp.Start(appState.CacheServerIP) }()
	go func() { proxy.Start(appState) }()

	if err := tea.NewProgram(tui.UIModel(details)).Start(); err != nil {
		log.Fatal(err)
	}
}
