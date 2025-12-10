package dns

import (
	"fmt"
	"log"
	"strings"

	"steam-lancache/internal/config"

	"github.com/miekg/dns"
)

var appState *config.AppState

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range m.Question {
			shouldSpoof := false
			for _, d := range config.AllDomains {
				if !strings.HasSuffix(d, ".") {
					d = d + "."
				}
				if strings.HasPrefix(d, "*.") {

					if strings.HasSuffix(q.Name, d[2:]) {
						shouldSpoof = true
						break
					}
				} else {
					if strings.HasSuffix(q.Name, d) {
						shouldSpoof = true
						break
					}
				}
			}

			if shouldSpoof {
				// spoof to local ip
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, config.LocalIP))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			} else {
				// forward to external dns
				c := new(dns.Client)
				upstream := "8.8.8.8:53"
				if appState.DNSIP != "" {
					upstream = appState.DNSIP + ":53"
				}
				in, _, err := c.Exchange(r, upstream)
				if err != nil {
					log.Printf("[DNS] Forward Error: %v", err)
					continue
				}
				m.Answer = append(m.Answer, in.Answer...)
			}
		}
	}
	w.WriteMsg(m)
}

func Start(state *config.AppState) {
	appState = state
	dnsServer := &dns.Server{Addr: ":" + config.DNSPort, Net: "udp"}
	dns.HandleFunc(".", handleDNSRequest)
	log.Printf("[DNS] Listening on UDP -> %s", config.DNSPort)
	if err := dnsServer.ListenAndServe(); err != nil {
		log.Printf("[DNS] Failed to start DNS server: %s", err.Error())
	}
}
