package stats

import (
	"steam-lancache/internal/config"
	"strings"
)

func DetectCategory(host string) string {
	h := strings.ToLower(host)

	for _, d := range config.Domains.Steam {
		if match(h, d) {
			return "steam"
		}
	}
	for _, d := range config.Domains.PlayStation {
		if match(h, d) {
			return "playstation"
		}
	}
	for _, d := range config.Domains.Xbox {
		if match(h, d) {
			return "xbox"
		}
	}
	for _, d := range config.Domains.Riot {
		if match(h, d) {
			return "riot"
		}
	}
	for _, d := range config.Domains.Epic {
		if match(h, d) {
			return "epic"
		}
	}
	for _, d := range config.Domains.Global {
		if match(h, d) {
			return "global"
		}
	}

	return "unknown"
}

func match(host, pattern string) bool {
	if strings.HasPrefix(pattern, "*.") {
		suffix := pattern[1:]
		return strings.HasSuffix(host, suffix)
	}
	return host == pattern
}
