package config

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	TokenAPI        = "https://api.steamdl.ir/get_user?token=%s"
	DNSListAPI      = "https://files.steamdl.ir/anti_sanction_dns.json"
	DomainListAPI   = "https://raw.githubusercontent.com/nimazerobit/steamdl-client/refs/heads/dev/domains.json"
	CacheDomain     = "dl.steamdl.ir"
	LocalIP         = "127.0.0.1"
	DNSPort         = "53"
	HTTPPort        = "80"
	HTTPSPort       = "443"
	StatsFile       = "usage.json"
	StatsUpdateFreq = 2 * time.Second
)

type DomainConfig struct {
	Global      []string `json:"global"`
	Steam       []string `json:"steam"`
	PlayStation []string `json:"playstation"`
	Xbox        []string `json:"xbox"`
	Riot        []string `json:"riot"`
	Epic        []string `json:"epic"`
}

var Domains DomainConfig

func LoadDomains() error {
	resp, err := http.Get(DomainListAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &Domains)
}

// runtime configuration
type AppState struct {
	UserToken     string
	CacheServerIP string
	DNSIP         string
}
