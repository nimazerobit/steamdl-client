package stats

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"steam-lancache/internal/config"
)

type CategoryStats struct {
	Bytes int64 `json:"bytes"`
}

var (
	mutex   sync.Mutex
	traffic = map[string]*CategoryStats{
		"steam":       {0},
		"playstation": {0},
		"xbox":        {0},
		"riot":        {0},
		"epic":        {0},
		"global":      {0},
	}
)

func Load() {
	data, err := os.ReadFile(config.StatsFile)
	if err != nil {
		return
	}

	_ = json.Unmarshal(data, &traffic)
}

func Add(category string, n int64) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := traffic[category]; !ok {
		traffic[category] = &CategoryStats{}
	}
	traffic[category].Bytes += n
}

func Snapshot() map[string]int64 {
	mutex.Lock()
	defer mutex.Unlock()

	out := make(map[string]int64)
	for k, v := range traffic {
		out[k] = v.Bytes
	}
	return out
}

func StartSaver() {
	ticker := time.NewTicker(config.StatsUpdateFreq)
	for range ticker.C {
		mutex.Lock()
		data, _ := json.MarshalIndent(traffic, "", "  ")
		mutex.Unlock()

		err := os.WriteFile(config.StatsFile, data, 0644)
		if err != nil {
			log.Printf("[Stats] Error saving: %v", err)
		}
	}
}
